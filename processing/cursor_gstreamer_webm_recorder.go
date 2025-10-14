package processing

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/GetStream/getstream-go/v3"
	"github.com/pion/rtp"
)

type CursorGstreamerWebmRecorder struct {
	logger          *getstream.DefaultLogger
	outputPath      string
	rtpConn         *net.UDPConn
	rtcpConn        *net.UDPConn
	gstreamerCmd    *exec.Cmd
	mu              sync.Mutex
	ctx             context.Context
	cancel          context.CancelFunc
	port            int
	sdpFile         *os.File
	finalOutputPath string // Path for post-processed file with duration
	tempOutputPath  string // Path for temporary file before post-processing
}

func NewCursorGstreamerWebmRecorder(outputPath, sdpContent string, logger *getstream.DefaultLogger) (*CursorGstreamerWebmRecorder, error) {
	ctx, cancel := context.WithCancel(context.Background())

	r := &CursorGstreamerWebmRecorder{
		logger:     logger,
		outputPath: outputPath,
		ctx:        ctx,
		cancel:     cancel,
	}

	// Set up UDP connections
	r.port = rand.Intn(10000) + 10000
	if err := r.setupConnections(r.port, true); err != nil {
		cancel()
		return nil, err
	}
	// Use rtcp on r.port+1 to match RTP/RTCP convention
	if err := r.setupConnections(r.port+1, false); err != nil {
		cancel()
		return nil, err
	}

	// Start GStreamer with codec detection
	if err := r.startGStreamer(sdpContent, outputPath); err != nil {
		cancel()
		return nil, err
	}

	return r, nil
}

func (r *CursorGstreamerWebmRecorder) setupConnections(port int, rtp bool) error {
	// Setup UDP connection
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	// Increase socket send buffer to reduce kernel-level drops
	_ = conn.SetWriteBuffer(4 << 20) // 4 MiB
	if rtp {
		r.rtpConn = conn
	} else {
		r.rtcpConn = conn
	}
	return nil
}

func (r *CursorGstreamerWebmRecorder) startGStreamer(sdpContent, outputFilePath string) error {
	// Write SDP to a temporary file
	sdpFile, err := os.CreateTemp("", "cursor_gstreamer_webm_*.sdp")
	if err != nil {
		return err
	}
	r.sdpFile = sdpFile

	updatedSdp := replaceSDP(sdpContent, r.port)

	if _, err := sdpFile.WriteString(updatedSdp); err != nil {
		sdpFile.Close()
		return err
	}
	sdpFile.Close()

	r.logger.Info("SDP created for GStreamer\n%s\n", updatedSdp)

	// Determine codec from SDP content and build GStreamer arguments
	isVP9 := strings.Contains(strings.ToUpper(sdpContent), "VP9")
	isVP8 := strings.Contains(strings.ToUpper(sdpContent), "VP8")
	isAV1 := strings.Contains(strings.ToUpper(sdpContent), "AV1")
	isH264 := strings.Contains(strings.ToUpper(sdpContent), "H264") || strings.Contains(strings.ToUpper(sdpContent), "H.264")
	isOpus := strings.Contains(strings.ToUpper(sdpContent), "OPUS")

	// Start with common GStreamer arguments optimized for RTP dump replay
	args := []string{
		"--gst-debug-level=2",
		//"--gst-debug=udpsrc:5,rtp*:5,webm*:5,identity:5,jitterbuffer:5,vp9*:5",
		//"--gst-debug-no-color",
		"-e", // Send EOS on interrupt for clean shutdown
	}
	// Add SDP source - this handles UDP connection and RTP setup automatically
	args = append(args,
		"sdpsrc",
		"location=sdp://"+sdpFile.Name(),
		"name=sdp",
		"sdp.stream_0",
		"!",
		// Add a large in-process queue to absorb bursts and decouple socket IO from depay
		"queue",
		"max-size-buffers=0",
		"max-size-bytes=268435456", // 256 MiB
		"max-size-time=0",
		"leaky=0",
		"!",
	)

	// Build pipeline based on codec with simplified RTP timestamp handling for dump replay
	//
	// Simplified approach for RTP dump replay:
	// - rtpjitterbuffer: Basic packet reordering with minimal interference
	//   - latency=0: No artificial latency, process packets as they come
	//   - mode=none: Don't override timing, let depayloaders handle it
	//   - do-retransmission=false: No retransmission for dump replay
	// - Remove identity sync to avoid timing conflicts
	//
	// This approach focuses on preserving original RTP timestamps without
	// artificial buffering that can interfere with dump replay timing.
	if false && isH264 {
		r.logger.Info("Detected H.264 codec, building H.264 pipeline with timestamp handling...")
		args = append(args,
			"application/x-rtp,media=video,encoding-name=H264,clock-rate=90000", "!",
			"rtpjitterbuffer",
			"latency=0",
			"mode=none",
			"do-retransmission=false", "!",
			"rtph264depay", "!",
			"h264parse", "!",
			"mp4mux", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	} else if false && isVP9 {
		r.logger.Info("Detected VP9 codec, building VP9 pipeline with timestamp handling...")
		args = append(args,
			"rtpjitterbuffer",
			"latency=0",
			"mode=none",
			"do-retransmission=false",
			"drop-on-latency=false",
			"buffer-mode=slave",
			"max-dropout-time=5000000000",
			"max-reorder-delay=1000000000",
			"!",
			"rtpvp9depay", "!",
			"vp9parse", "!",
			"webmmux",
			"writing-app=GStreamer-VP9",
			"streamable=false",
			"min-index-interval=2000000000", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	} else if isVP9 {
		r.logger.Info("Detected VP9 codec, building VP9 pipeline with RTP timestamp handling...")
		args = append(args,

			//// jitterbuffer for packet reordering and timestamp handling
			"rtpjitterbuffer",
			"name=jitterbuffer",
			"mode=none",
			"latency=0",               // No artificial latency - process immediately
			"do-lost=false",           // Don't generate lost events for missing packets
			"do-retransmission=false", // No retransmission for offline replay
			"drop-on-latency=false",   // Keep all packets even if late
			"!",
			//
			// Depayload RTP to get VP9 frames
			"rtpvp9depay",
			"!",

			// Parse VP9 stream to ensure valid frame structure
			"vp9parse",
			"!",

			// Queue for buffering
			"queue",
			"!",

			// Mux into Matroska/WebM container
			"webmmux",
			"writing-app=GStreamer-VP9",
			"streamable=false",
			"min-index-interval=2000000000",
			"!",

			// Write to file
			"filesink",
			fmt.Sprintf("location=%s", outputFilePath),
		)

	} else if false && isVP8 {
		r.logger.Info("Detected VP8 codec, building VP8 pipeline with timestamp handling...")
		args = append(args,
			"application/x-rtp,media=video,encoding-name=VP8,clock-rate=90000", "!",
			"rtpjitterbuffer",
			"latency=0",
			"mode=none",
			"do-retransmission=false", "!",
			"rtpvp8depay", "!",
			"vp8parse", "!",
			"webmmux", "writing-app=GStreamer", "streamable=false", "min-index-interval=2000000000", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	} else if false && isAV1 {
		r.logger.Info("Detected AV1 codec, building AV1 pipeline with timestamp handling...")
		args = append(args,
			"application/x-rtp,media=video,encoding-name=AV1,clock-rate=90000", "!",
			"rtpjitterbuffer",
			"latency=0",
			"mode=none",
			"do-retransmission=false", "!",
			"rtpav1depay", "!",
			"av1parse", "!",
			"webmmux", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	} else if false && isOpus {
		r.logger.Info("Detected Opus codec, building Opus pipeline with timestamp handling...")
		args = append(args,
			"application/x-rtp,media=audio,encoding-name=OPUS,clock-rate=48000,payload=111", "!",
			"rtpjitterbuffer",
			"latency=0",
			"mode=none",
			"do-retransmission=false", "!",
			"rtpopusdepay", "!",
			"opusparse", "!",
			"webmmux", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	} else if false {
		// Default to VP8 if codec is not detected
		r.logger.Info("Unknown or no codec detected, defaulting to VP8 pipeline with timestamp handling...")
		args = append(args,
			"application/x-rtp,media=video,encoding-name=VP8,clock-rate=90000", "!",
			"rtpjitterbuffer",
			"latency=0",
			"mode=none",
			"do-retransmission=false", "!",
			"rtpvp8depay", "!",
			"vp8parse", "!",
			"webmmux", "writing-app=GStreamer", "streamable=false", "min-index-interval=2000000000", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	}

	r.logger.Info("GStreamer pipeline: %s", strings.Join(args, " ")) // Skip debug args for display

	r.gstreamerCmd = exec.Command("gst-launch-1.0", args...)

	// Redirect output for debugging
	r.gstreamerCmd.Stdout = os.Stdout
	r.gstreamerCmd.Stderr = os.Stderr

	// Start GStreamer process
	if err := r.gstreamerCmd.Start(); err != nil {
		return err
	}

	r.logger.Info("GStreamer pipeline started with PID: %d", r.gstreamerCmd.Process.Pid)

	// Monitor the process in a goroutine
	go func() {
		if err := r.gstreamerCmd.Wait(); err != nil {
			r.logger.Error("GStreamer process exited with error: %v", err)
		} else {
			r.logger.Info("GStreamer process exited normally")
		}
	}()

	return nil
}

func (r *CursorGstreamerWebmRecorder) OnRTP(packet *rtp.Packet) error {
	// Marshal RTP packet
	buf, err := packet.Marshal()
	if err != nil {
		return err
	}

	return r.PushRtpBuf(buf)
}

func (r *CursorGstreamerWebmRecorder) OnRTCP(packet *rtp.Packet) error {
	// Marshal RTP packet
	buf, err := packet.Marshal()
	if err != nil {
		return err
	}

	return r.PushRtcpBuf(buf)
}

func (r *CursorGstreamerWebmRecorder) PushRtcpBuf(buf []byte) error {
	return r.pushBuf(r.rtcpConn, buf)
}

func (r *CursorGstreamerWebmRecorder) PushRtpBuf(buf []byte) error {
	return r.pushBuf(r.rtpConn, buf)
}

func (r *CursorGstreamerWebmRecorder) pushBuf(conn *net.UDPConn, buf []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Send RTP packet over UDP to GStreamer udpsrc
	if conn != nil {
		_, err := conn.Write(buf)
		if err != nil {
			// Log error but don't fail completely - some packet loss is acceptable
			r.logger.Warn("Failed to write RTP packet: %v", err)
		}
	}
	return nil
}

func (r *CursorGstreamerWebmRecorder) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Info("Closing GStreamer WebM recorder...")

	r.logger.Info("Closing GStreamer WebM recorder2222...")

	// Cancel context to stop background goroutines
	if r.cancel != nil {
		r.cancel()
	}

	// Close UDP connection with goodbye message
	if r.rtpConn != nil {
		r.logger.Info("Closing UDP connection...")

		// Send RTCP Goodbye packet to signal end of stream
		//buf, _ := rtcp.Goodbye{
		//	Sources: []uint32{1}, // fixed SSRC is ok
		//	Reason:  "bye",
		//}.Marshal()
		//_, _ = r.rtpConn.Write(buf)

		r.logger.Info("Goodbye sent")

		// Give some time for the goodbye packet to be processed
		time.Sleep(1 * time.Second)

		_ = r.rtpConn.Close()
		r.rtpConn = nil
		r.logger.Info("UDP connection closed")
	}

	// Gracefully stop GStreamer
	if r.gstreamerCmd != nil && r.gstreamerCmd.Process != nil {
		r.logger.Info("Stopping GStreamer process...")

		// Send EOS (End of Stream) signal to GStreamer
		// GStreamer handles SIGINT gracefully and will finish writing the file
		if err := r.gstreamerCmd.Process.Signal(os.Interrupt); err != nil {
			r.logger.Error("Failed to send SIGINT to GStreamer: %v", err)
			// If interrupt fails, force kill
			r.gstreamerCmd.Process.Kill()
		} else {
			r.logger.Info("Sent SIGINT to GStreamer, waiting for graceful exit...")

			// Wait for graceful exit with timeout
			done := make(chan error, 1)
			go func() {
				done <- r.gstreamerCmd.Wait()
			}()

			select {
			case <-time.After(15 * time.Second):
				r.logger.Info("GStreamer exit timeout, force killing...")
				// Timeout, force kill
				r.gstreamerCmd.Process.Kill()
				<-done // Wait for the kill to complete
			case err := <-done:
				if err != nil {
					r.logger.Info("GStreamer exited with error: %v", err)
				} else {
					r.logger.Info("GStreamer exited gracefully")
				}
			}
		}
	}

	// Clean up temporary SDP file
	if r.sdpFile != nil {
		os.Remove(r.sdpFile.Name())
		r.sdpFile = nil
	}

	// Post-process WebM to fix duration metadata if needed
	if r.tempOutputPath != "" && r.finalOutputPath != "" {
		r.logger.Info("Starting WebM duration post-processing...")
	}

	r.logger.Info("GStreamer WebM recorder closed")
	return nil
}

// GetOutputPath returns the output file path (for compatibility)
func (r *CursorGstreamerWebmRecorder) GetOutputPath() string {
	// Return final output path if post-processing is enabled, otherwise return original
	if r.finalOutputPath != "" {
		return r.finalOutputPath
	}
	return r.outputPath
}

// IsRecording returns true if the recorder is currently active
func (r *CursorGstreamerWebmRecorder) IsRecording() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.gstreamerCmd != nil && r.gstreamerCmd.Process != nil
}
