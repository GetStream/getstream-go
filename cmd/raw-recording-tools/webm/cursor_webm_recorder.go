package webm

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/GetStream/getstream-go/v3"
	"github.com/GetStream/getstream-go/v3/cmd/raw-recording-tools/rawsdputil"
	"github.com/pion/rtcp"
	"github.com/pion/rtp"
)

type CursorWebmRecorder struct {
	logger     *getstream.DefaultLogger
	outputPath string
	conn       *net.UDPConn
	ffmpegCmd  *exec.Cmd
	stdin      io.WriteCloser
	mu         sync.Mutex
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewCursorWebmRecorder(outputPath, sdpContent string, logger *getstream.DefaultLogger) (*CursorWebmRecorder, error) {
	ctx, cancel := context.WithCancel(context.Background())

	r := &CursorWebmRecorder{
		logger:     logger,
		outputPath: outputPath,
		ctx:        ctx,
		cancel:     cancel,
	}

	r.logger.Info("Sdp created \n%s\n", sdpContent)

	// Set up UDP connections
	port := rand.Intn(10000) + 10000
	if err := r.setupConnections(port); err != nil {
		cancel()
		return nil, err
	}

	// Start FFmpeg with codec detection
	if err := r.startFFmpeg(outputPath, sdpContent, port); err != nil {
		cancel()
		return nil, err
	}

	return r, nil
}

func (r *CursorWebmRecorder) setupConnections(port int) error {
	// Setup connection
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

func (r *CursorWebmRecorder) startFFmpeg(outputFilePath, sdpContent string, port int) error {

	// Write SDP to a temporary file
	sdpFile, err := os.CreateTemp("", "cursor_webm_*.sdp")
	if err != nil {
		return err
	}

	if _, err := sdpFile.WriteString(rawsdputil.ReplaceSDP(sdpContent, port)); err != nil {
		sdpFile.Close()
		return err
	}
	sdpFile.Close()

	// Build FFmpeg command with optimized settings for single track recording
	args := []string{
		"-threads", "1",
		//		"-loglevel", "debug",
		"-protocol_whitelist", "file,udp,rtp",
		"-buffer_size", "10000000",
		"-max_delay", "150000",
		"-reorder_queue_size", "130",
		"-i", sdpFile.Name(),
	}

	//switch strings.ToLower(mimeType) {
	//case "audio/opus":
	//	// For other codecs, use direct copy
	args = append(args, "-c", "copy")
	//default:
	//	// For other codecs, use direct copy
	//	args = append(args, "-c", "copy")
	//}
	//if isVP9 {
	//	// For VP9, avoid direct copy and use re-encoding with error resilience
	//	// This works around FFmpeg's experimental VP9 RTP support issues
	//	r.logger.Info("Detected VP9 codec, applying workarounds...")
	//	args = append(args,
	//		"-c:v", "libvpx-vp9",
	//		//			"-error_resilience", "aggressive",
	//		"-err_detect", "ignore_err",
	//		"-fflags", "+genpts+igndts",
	//		"-avoid_negative_ts", "make_zero",
	//		// VP9-specific quality settings to handle corrupted frames
	//		"-crf", "30",
	//		"-row-mt", "1",
	//		"-frame-parallel", "1",
	//	)
	//} else if strings.Contains(strings.ToUpper(sdpContent), "AV1") {
	//	args = append(args,
	//		"-c:v", "libaom-av1",
	//		"-cpu-used", "8",
	//		"-usage", "realtime",
	//	)
	//} else if strings.Contains(strings.ToUpper(sdpContent), "OPUS") {
	//	args = append(args, "-fflags", "+genpts", "-use_wallclock_as_timestamps", "0", "-c:a", "copy")
	//} else {
	//	// For other codecs, use direct copy
	//	args = append(args, "-c", "copy")
	//}

	args = append(args,
		"-y",
		outputFilePath,
	)

	r.ffmpegCmd = exec.Command("ffmpeg", args...)

	// Redirect output for debugging
	r.ffmpegCmd.Stdout = os.Stdout
	r.ffmpegCmd.Stderr = os.Stderr

	// Create stdin pipe to send commands to FFmpeg
	//var err error
	r.stdin, err = r.ffmpegCmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe:", err)
	}

	// Start FFmpeg process
	if err := r.ffmpegCmd.Start(); err != nil {
		return err
	}

	return nil
}

func (r *CursorWebmRecorder) OnRTP(packet *rtp.Packet) error {
	// Marshal RTP packet
	buf, err := packet.Marshal()
	if err != nil {
		return err
	}

	return r.PushRtpBuf(buf)
}

func (r *CursorWebmRecorder) PushRtpBuf(buf []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Send RTP packet over UDP
	if r.conn != nil {
		_, _ = r.conn.Write(buf)
		//if err != nil {
		//	return err)
		//}
		//		r.logger.Info("Wrote packet to %s - %v", r.conn.LocalAddr().String(), err)
	}
	return nil
}

func (r *CursorWebmRecorder) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Cancel context to stop background goroutines
	if r.cancel != nil {
		r.cancel()
	}

	r.logger.Info("Closing UPD connection...")

	// Close UDP connection by sending arbitrary RtcpBye (Ffmpeg is no able to end correctly)
	if r.conn != nil {
		buf, _ := rtcp.Goodbye{
			Sources: []uint32{1}, // fixed ssrc is ok
			Reason:  "bye",
		}.Marshal()
		_, _ = r.conn.Write(buf)
		_ = r.conn.Close()
		r.conn = nil
	}

	r.logger.Info("UDP Connection closed...")

	time.Sleep(5 * time.Second)

	r.logger.Info("After sleep...")

	// Gracefully stop FFmpeg
	if r.ffmpegCmd != nil && r.ffmpegCmd.Process != nil {

		// ✅ Gracefully stop FFmpeg by sending 'q' to stdin
		//fmt.Println("Sending 'q' to FFmpeg...")
		//_, _ = r.stdin.Write([]byte("q\n"))
		//r.stdin.Close()

		// Send interrupt signal to FFmpeg process
		r.logger.Info("Sending SIGTERM...")

		//if err := r.ffmpegCmd.Process.Signal(os.Interrupt); err != nil {
		//	// If interrupt fails, force kill
		//	r.ffmpegCmd.Process.Kill()
		//} else {

		r.logger.Info("Waiting for SIGTERM...")

		// Wait for graceful exit with timeout
		done := make(chan error, 1)
		go func() {
			done <- r.ffmpegCmd.Wait()
		}()

		select {
		case <-time.After(10 * time.Second):
			r.logger.Info("Wait timetout for SIGTERM...")

			// Timeout, force kill
			r.ffmpegCmd.Process.Kill()
		case <-done:
			r.logger.Info("Process exited succesfully SIGTERM...")
			// Process exited gracefully
		}
	}

	return nil
}
