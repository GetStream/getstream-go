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
	conn            *net.UDPConn
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

	r.logger.Info("SDP created for GStreamer\n%s\n", sdpContent)

	// Set up UDP connections
	r.port = rand.Intn(10000) + 10000
	if err := r.setupConnections(r.port); err != nil {
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

func (r *CursorGstreamerWebmRecorder) setupConnections(port int) error {
	// Setup UDP connection
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	r.conn = conn

	return nil
}

func (r *CursorGstreamerWebmRecorder) startGStreamer(sdpContent, outputFilePath string) error {
	// Write SDP to a temporary file
	sdpFile, err := os.CreateTemp("", "cursor_gstreamer_webm_*.sdp")
	if err != nil {
		return err
	}
	r.sdpFile = sdpFile

	if _, err := sdpFile.WriteString(sdpContent); err != nil {
		sdpFile.Close()
		return err
	}
	sdpFile.Close()

	// Determine codec from SDP content and build GStreamer arguments
	isVP9 := strings.Contains(strings.ToUpper(sdpContent), "VP9")
	isVP8 := strings.Contains(strings.ToUpper(sdpContent), "VP8")
	isAV1 := strings.Contains(strings.ToUpper(sdpContent), "AV1")
	isH264 := strings.Contains(strings.ToUpper(sdpContent), "H264") || strings.Contains(strings.ToUpper(sdpContent), "H.264")
	isOpus := strings.Contains(strings.ToUpper(sdpContent), "OPUS")

	// Start with common GStreamer arguments optimized for RTP dump replay
	args := []string{
		"--gst-debug-level=3",
		"--gst-debug=udpsrc:5,rtp*:5,webm*:5,identity:5,jitterbuffer:5",
	}

	// Add UDP source with timestamp handling for RTP dump replay
	args = append(args,
		"-e",
		"udpsrc",
		fmt.Sprintf("port=%d", r.port),
		"buffer-size=10000000",
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
	if isH264 {
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
	} else if isVP9 {
		r.logger.Info("Detected VP9 codec, building VP9 pipeline with timestamp handling...")
		args = append(args,
			"application/x-rtp,media=video,encoding-name=VP9,clock-rate=90000", "!",
			"rtpjitterbuffer",
			"latency=200",
			"mode=none",
			"do-retransmission=false", "!",
			"rtpvp9depay", "!",
			"vp9parse", "!",
			"webmmux", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	} else if isVP8 {
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
	} else if isAV1 {
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
	} else if isOpus {
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
	} else {
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

	r.logger.Info("GStreamer pipeline: %s", strings.Join(args[3:], " ")) // Skip debug args for display

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

func (r *CursorGstreamerWebmRecorder) PushRtpBuf(buf []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Send RTP packet over UDP to GStreamer udpsrc
	if r.conn != nil {
		_, err := r.conn.Write(buf)
		if err != nil {
			// Log error but don't fail completely - some packet loss is acceptable
			r.logger.Debug("Failed to write RTP packet: %v", err)
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
	if r.conn != nil {
		r.logger.Info("Closing UDP connection...")

		// Send RTCP Goodbye packet to signal end of stream
		//buf, _ := rtcp.Goodbye{
		//	Sources: []uint32{1}, // fixed SSRC is ok
		//	Reason:  "bye",
		//}.Marshal()
		//_, _ = r.conn.Write(buf)

		r.logger.Info("Goodbye sent")

		// Give some time for the goodbye packet to be processed
		time.Sleep(1 * time.Second)

		_ = r.conn.Close()
		r.conn = nil
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

		// Choose post-processing approach based on temp file extension
		if strings.HasSuffix(r.tempOutputPath, ".gst") {
			// Simple approach for .gst files
			if err := r.simpleWebMDurationFix(); err != nil {
				r.logger.Error("Simple WebM duration fix failed: %v", err)
			}
		} else if strings.HasSuffix(r.tempOutputPath, ".direct") {
			// Simple approach for .direct files (direct timing with post-processing)
			if err := r.simpleWebMDurationFix(); err != nil {
				r.logger.Error("Direct WebM duration fix failed: %v", err)
			}
		} else if strings.HasSuffix(r.tempOutputPath, ".minimal") {
			// Simple approach for .minimal files (minimal timing with post-processing)
			if err := r.simpleWebMDurationFix(); err != nil {
				r.logger.Error("Minimal WebM duration fix failed: %v", err)
			}
		} else {
			// Enhanced approach for .temp files
			if err := r.postProcessWebMDuration(); err != nil {
				r.logger.Error("Enhanced WebM post-processing failed: %v", err)
			}
		}
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

// BufferConfig holds the configuration for RTP jitter buffer settings
type BufferConfig struct {
	Latency         int  // Buffer latency in milliseconds
	MaxMisorderTime int  // Maximum time to wait for out-of-order packets (ms)
	MaxDropoutTime  int  // Maximum time before considering packet lost (ms)
	RtxDelay        int  // Retransmission delay in milliseconds
	DoLost          bool // Generate lost packet events
	DropOnLatency   bool // Drop packets that arrive too late
}

// DefaultBufferConfig returns optimized settings for RTP dump replay with reordering
func DefaultBufferConfig() BufferConfig {
	return BufferConfig{
		Latency:         500,   // 500ms buffer for reordering
		MaxMisorderTime: 2000,  // Wait up to 2 seconds for missing packets
		MaxDropoutTime:  60000, // Consider packets lost after 60 seconds
		RtxDelay:        40,    // Request retransmission after 40ms
		DoLost:          true,  // Generate lost packet events for debugging
		DropOnLatency:   false, // Don't drop packets, buffer them for proper ordering
	}
}

// RealtimeBufferConfig returns optimized settings for real-time streaming
func RealtimeBufferConfig() BufferConfig {
	return BufferConfig{
		Latency:         100,  // Lower latency for real-time
		MaxMisorderTime: 500,  // Shorter wait time
		MaxDropoutTime:  5000, // Faster dropout detection
		RtxDelay:        20,   // Faster retransmission
		DoLost:          true,
		DropOnLatency:   true, // Drop late packets to maintain real-time performance
	}
}

// NewCursorGstreamerWebmRecorderNoJitterBuffer creates a recorder without jitter buffer for direct timing
// This bypasses all buffering and lets the depayloaders handle RTP timestamps directly.
func NewCursorGstreamerWebmRecorderNoJitterBuffer(outputPath, sdp string, port int) (*CursorGstreamerWebmRecorder, error) {
	ctx, cancel := context.WithCancel(context.Background())

	r := &CursorGstreamerWebmRecorder{
		outputPath: outputPath,
		ctx:        ctx,
		cancel:     cancel,
		port:       port,
	}

	r.logger.Info("SDP created for GStreamer (no jitter buffer)\n%s\n", sdp)

	// Set up UDP connections
	if err := r.setupConnections(port); err != nil {
		cancel()
		return nil, err
	}

	// Start GStreamer without jitter buffer
	if err := r.startGStreamerNoJitterBuffer(sdp, outputPath); err != nil {
		cancel()
		return nil, err
	}

	return r, nil
}

func (r *CursorGstreamerWebmRecorder) startGStreamerNoJitterBuffer(sdpContent, outputFilePath string) error {
	// Write SDP to a temporary file
	sdpFile, err := os.CreateTemp("", "cursor_gstreamer_webm_*.sdp")
	if err != nil {
		return err
	}
	r.sdpFile = sdpFile

	if _, err := sdpFile.WriteString(sdpContent); err != nil {
		sdpFile.Close()
		return err
	}
	sdpFile.Close()

	// Determine codec from SDP content
	isVP9 := strings.Contains(strings.ToUpper(sdpContent), "VP9")
	isVP8 := strings.Contains(strings.ToUpper(sdpContent), "VP8")
	isAV1 := strings.Contains(strings.ToUpper(sdpContent), "AV1")
	isH264 := strings.Contains(strings.ToUpper(sdpContent), "H264") || strings.Contains(strings.ToUpper(sdpContent), "H.264")

	// Start with common GStreamer arguments
	args := []string{
		"--gst-debug-level=3",
		"--gst-debug=udpsrc:5,rtp*:5,webm*:5",
		"-e", // Enable EOS handling
	}

	// Add UDP source with timing preservation for direct recording
	args = append(args,
		"udpsrc",
		fmt.Sprintf("port=%d", r.port),
		"buffer-size=10000000",
		"!",
		"queue",
		"max-size-buffers=1000",
		"max-size-time=10000000000", // 10 seconds of buffering
		"!",
	)

	// Build pipeline based on codec WITHOUT jitter buffer
	// Use provided output file path directly
	if isH264 {
		r.logger.Info("Detected H.264 codec, building direct H.264 pipeline...")
		args = append(args,
			"application/x-rtp,media=video,encoding-name=H264,clock-rate=90000", "!",
			"rtph264depay", "!",
			"h264parse", "!",
			"identity", "sync=true", "!", // Force timing synchronization
			"mp4mux", "faststart=true", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	} else if isVP9 {
		r.logger.Info("Detected VP9 codec, building direct VP9 pipeline...")
		args = append(args,
			"application/x-rtp,media=video,encoding-name=VP9,clock-rate=90000", "!",
			"rtpvp9depay", "!",
			"vp9parse", "!",
			"identity", "sync=true", "!", // Force timing synchronization
			"webmmux", "writing-app=GStreamer-Direct", "streamable=false", "min-index-interval=1000000000", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	} else if isVP8 {
		r.logger.Info("Detected VP8 codec, building direct VP8 pipeline...")
		args = append(args,
			"application/x-rtp,media=video,encoding-name=VP8,clock-rate=90000", "!",
			"rtpvp8depay", "!",
			"vp8parse", "!",
			"identity", "sync=true", "!", // Force timing synchronization
			"webmmux", "writing-app=GStreamer-Direct", "streamable=false", "min-index-interval=1000000000", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	} else if isAV1 {
		r.logger.Info("Detected AV1 codec, building direct AV1 pipeline...")
		args = append(args,
			"application/x-rtp,media=video,encoding-name=AV1,clock-rate=90000", "!",
			"rtpav1depay", "!",
			"av1parse", "!",
			"identity", "sync=true", "!", // Force timing synchronization
			"webmmux", "writing-app=GStreamer-Direct", "streamable=false", "min-index-interval=1000000000", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	} else {
		// Default to VP8 if codec is not detected
		r.logger.Info("Unknown or no codec detected, defaulting to direct VP8 pipeline...")
		args = append(args,
			"application/x-rtp,media=video,encoding-name=VP8,clock-rate=90000", "!",
			"rtpvp8depay", "!",
			"vp8parse", "!",
			"identity", "sync=true", "!", // Force timing synchronization
			"webmmux", "writing-app=GStreamer-Direct", "streamable=false", "min-index-interval=1000000000", "!",
			"filesink", fmt.Sprintf("location=%s", outputFilePath),
		)
	}

	r.logger.Info("GStreamer direct pipeline: %s", strings.Join(args[3:], " "))

	r.gstreamerCmd = exec.CommandContext(r.ctx, "gst-launch-1.0", args...)

	// Redirect output for debugging
	r.gstreamerCmd.Stdout = os.Stdout
	r.gstreamerCmd.Stderr = os.Stderr

	// Start GStreamer process
	if err := r.gstreamerCmd.Start(); err != nil {
		return err
	}

	r.logger.Info("GStreamer direct pipeline started with PID: %d", r.gstreamerCmd.Process.Pid)

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

// postProcessWebMDuration fixes WebM duration metadata using FFmpeg
// This ensures the WebM file has proper duration information for browser playback
func (r *CursorGstreamerWebmRecorder) postProcessWebMDuration() error {
	if r.tempOutputPath == "" || r.finalOutputPath == "" {
		// No post-processing needed
		return nil
	}

	r.logger.Info("Post-processing WebM duration metadata...")

	// Check if temp file exists
	if _, err := os.Stat(r.tempOutputPath); os.IsNotExist(err) {
		r.logger.Warn("Temp file does not exist for post-processing: %s", r.tempOutputPath)
		return nil
	}

	// First get the duration from the file
	durationCmd := exec.Command("ffprobe",
		"-v", "quiet",
		"-show_entries", "format=duration",
		"-of", "csv=p=0",
		r.tempOutputPath,
	)

	durationOutput, err := durationCmd.Output()
	if err != nil {
		r.logger.Error("Failed to get duration with ffprobe: %v", err)
	}

	duration := strings.TrimSpace(string(durationOutput))
	r.logger.Info("Detected file duration: %s seconds", duration)

	// Use more aggressive FFmpeg approach to ensure duration is written to WebM header
	cmd := exec.Command("ffmpeg",
		"-i", r.tempOutputPath,
		"-c:v", "copy", // Copy video stream
		"-avoid_negative_ts", "make_zero",
		"-fflags", "+genpts", // Generate presentation timestamps
		"-f", "webm", // Force WebM format
		"-write_crc32", "0", // Disable CRC for compatibility
		"-cluster_size_limit", "2097152", // 2MB clusters for better seeking
		"-cluster_time_limit", "5000", // 5 second clusters
		"-y", // Overwrite output file
		r.finalOutputPath,
	)

	// Set up logging
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	r.logger.Info("Running enhanced FFmpeg post-processing: %s", strings.Join(cmd.Args, " "))

	// Run FFmpeg
	if err := cmd.Run(); err != nil {
		r.logger.Error("Enhanced FFmpeg post-processing failed: %v", err)

		// Try a simpler WebM remux approach
		r.logger.Info("Trying fallback WebM remux...")
		fallbackCmd := exec.Command("ffmpeg",
			"-i", r.tempOutputPath,
			"-c", "copy",
			"-f", "webm",
			"-y",
			r.finalOutputPath,
		)

		fallbackCmd.Stdout = os.Stdout
		fallbackCmd.Stderr = os.Stderr

		if fallbackErr := fallbackCmd.Run(); fallbackErr != nil {
			r.logger.Error("Fallback FFmpeg also failed: %v", fallbackErr)
			// Last resort - just move the file
			return os.Rename(r.tempOutputPath, r.finalOutputPath)
		}
	}

	// Remove temporary file
	os.Remove(r.tempOutputPath)

	r.logger.Info("WebM duration metadata fixed successfully")
	return nil
}

// NewCursorGstreamerWebmRecorderWithDurationFix creates a recorder that automatically fixes WebM duration
// This version writes to a temporary file and post-processes it to ensure proper duration metadata
func NewCursorGstreamerWebmRecorderWithDurationFix(outputPath, sdp string, port int) (*CursorGstreamerWebmRecorder, error) {
	ctx, cancel := context.WithCancel(context.Background())

	r := &CursorGstreamerWebmRecorder{
		outputPath:      outputPath + ".temp", // Write to temp file first
		ctx:             ctx,
		cancel:          cancel,
		port:            port,
		finalOutputPath: outputPath,
		tempOutputPath:  outputPath + ".temp",
	}

	r.logger.Info("SDP created for GStreamer with duration fix\n%s\n", sdp)

	// Set up UDP connections
	if err := r.setupConnections(port); err != nil {
		cancel()
		return nil, err
	}

	// Start GStreamer
	if err := r.startGStreamer(sdp, r.tempOutputPath); err != nil {
		cancel()
		return nil, err
	}

	return r, nil
}

// NewCursorGstreamerWebmRecorderSimpleDuration creates a recorder with the simplest duration fix
// This version uses a minimal FFmpeg remux specifically for WebM duration metadata
func NewCursorGstreamerWebmRecorderSimpleDuration(outputPath, sdp string, port int) (*CursorGstreamerWebmRecorder, error) {
	ctx, cancel := context.WithCancel(context.Background())

	r := &CursorGstreamerWebmRecorder{
		outputPath:      outputPath + ".gst",
		ctx:             ctx,
		cancel:          cancel,
		port:            port,
		finalOutputPath: outputPath,
		tempOutputPath:  outputPath + ".gst",
	}

	r.logger.Info("Creating simple duration fix recorder\n%s\n", sdp)

	// Set up UDP connections
	if err := r.setupConnections(port); err != nil {
		cancel()
		return nil, err
	}

	// Start GStreamer
	if err := r.startGStreamer(sdp, r.tempOutputPath); err != nil {
		cancel()
		return nil, err
	}

	return r, nil
}

// simpleWebMDurationFix performs a minimal FFmpeg remux to fix duration for browsers
func (r *CursorGstreamerWebmRecorder) simpleWebMDurationFix() error {
	if r.tempOutputPath == "" || r.finalOutputPath == "" {
		return nil
	}

	r.logger.Info("Applying simple WebM duration fix...")

	// Check if temp file exists
	if _, err := os.Stat(r.tempOutputPath); os.IsNotExist(err) {
		r.logger.Warn("Source file does not exist: %s", r.tempOutputPath)
		return nil
	}

	// Simple FFmpeg remux that should preserve duration for browser playback
	cmd := exec.Command("ffmpeg",
		"-i", r.tempOutputPath,
		"-c", "copy", // Copy all streams
		"-f", "webm", // Ensure WebM format
		"-avoid_negative_ts", "make_zero",
		"-y",
		r.finalOutputPath,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	r.logger.Info("Running simple WebM fix: %s", strings.Join(cmd.Args, " "))

	if err := cmd.Run(); err != nil {
		r.logger.Error("Simple WebM fix failed: %v", err)
		// Fall back to just moving the file
		return os.Rename(r.tempOutputPath, r.finalOutputPath)
	}

	// Remove temp file
	os.Remove(r.tempOutputPath)

	r.logger.Info("Simple WebM duration fix completed")
	return nil
}
