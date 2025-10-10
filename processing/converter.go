package processing

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/GetStream/getstream-go/v3"
	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media/rtpdump"
	"github.com/pion/webrtc/v4/pkg/media/samplebuilder"
)

const audioMaxLate = 200  // 4sec
const videoMaxLate = 1000 // 4sec

type RTPDump2WebMConverter struct {
	logger        *getstream.DefaultLogger
	reader        *rtpdump.Reader
	recorder      WebmRecorder
	sampleBuilder *samplebuilder.SampleBuilder

	lastPkt         *rtp.Packet
	lastPktDuration uint32
	inserted        uint16
}

type WebmRecorder interface {
	OnRTP(pkt *rtp.Packet) error
	PushRtpBuf(payload []byte) error
	Close() error
}

func newRTPDump2WebMConverter(logger *getstream.DefaultLogger) *RTPDump2WebMConverter {
	return &RTPDump2WebMConverter{
		logger: logger,
	}
}

func ConvertDirectory(directory string, accept func(path string, info os.FileInfo) bool, fixDtx bool, logger *getstream.DefaultLogger) error {
	var rtpdumpFiles []string

	// Walk through directory to find .rtpdump files
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), SuffixRtpDump) && accept(path, info) {
			rtpdumpFiles = append(rtpdumpFiles, path)
		}

		return nil
	})
	if err != nil {
		return err
	}

	for _, rtpdumpFile := range rtpdumpFiles {
		c := newRTPDump2WebMConverter(logger)
		if err := c.ConvertFile(rtpdumpFile, fixDtx); err != nil {
			c.logger.Error("Failed to convert %s: %v", rtpdumpFile, err)
			continue
		}
	}

	return nil
}

func (c *RTPDump2WebMConverter) ConvertFile(inputFile string, fixDtx bool) error {
	c.logger.Info("Converting %s", inputFile)

	// Parse the RTP dump file
	// Open the file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open rtpdump file: %w", err)
	}
	defer file.Close()

	// Create standardized reader
	reader, _, _ := rtpdump.NewReader(file)
	c.reader = reader

	sdpContent, _ := readSDP(strings.Replace(inputFile, SuffixRtpDump, SuffixSdp, 1))
	mType, _ := mimeType(sdpContent)

	releasePacketHandler := samplebuilder.WithPacketReleaseHandler(c.buildDefaultReleasePacketHandler())

	switch mType {
	case webrtc.MimeTypeAV1:
		c.sampleBuilder = samplebuilder.New(videoMaxLate, &codecs.AV1Depacketizer{}, 90000, releasePacketHandler)
		c.recorder, err = NewCursorGstreamerWebmRecorder(strings.Replace(inputFile, SuffixRtpDump, SuffixWebm, 1), sdpContent, c.logger)
	case webrtc.MimeTypeVP9:
		c.sampleBuilder = samplebuilder.New(videoMaxLate, &codecs.VP9Packet{}, 90000, releasePacketHandler)
		c.recorder, err = NewCursorGstreamerWebmRecorder(strings.Replace(inputFile, SuffixRtpDump, SuffixWebm, 1), sdpContent, c.logger)
	case webrtc.MimeTypeH264:
		c.sampleBuilder = samplebuilder.New(videoMaxLate, &codecs.H264Packet{}, 90000, releasePacketHandler)
		c.recorder, err = NewCursorWebmRecorder(strings.Replace(inputFile, SuffixRtpDump, SuffixMp4, 1), sdpContent, c.logger)
	case webrtc.MimeTypeVP8:
		c.sampleBuilder = samplebuilder.New(videoMaxLate, &codecs.VP8Packet{}, 90000, releasePacketHandler)
		c.recorder, err = NewCursorWebmRecorder(strings.Replace(inputFile, SuffixRtpDump, SuffixWebm, 1), sdpContent, c.logger)
	case webrtc.MimeTypeOpus:
		if fixDtx {
			releasePacketHandler = samplebuilder.WithPacketReleaseHandler(c.buildOpusReleasePacketHandler())
		}
		c.sampleBuilder = samplebuilder.New(audioMaxLate, &codecs.OpusPacket{}, 48000, releasePacketHandler)
		c.recorder, err = NewCursorWebmRecorder(strings.Replace(inputFile, SuffixRtpDump, SuffixWebm, 1), sdpContent, c.logger)
	default:
		return fmt.Errorf("unsupported codec type: %s", mType)
	}
	if err != nil {
		return fmt.Errorf("failed to create WebM recorder: %w", err)
	}
	defer c.recorder.Close()

	time.Sleep(1 * time.Second)

	// Convert and feed RTP packets
	return c.feedPackets(reader)
}

func (c *RTPDump2WebMConverter) feedPackets(reader *rtpdump.Reader) error {
	startTime := time.Now()

	i := 0
	for ; ; i++ {
		packet, err := reader.Next()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		} else if packet.IsRTCP {
			continue
		}

		// Unmarshal the RTP packet from the raw payload
		if c.sampleBuilder == nil {
			_ = c.recorder.PushRtpBuf(packet.Payload)
		} else {
			// Unmarshal the RTP packet from the raw payload
			rtpPacket := &rtp.Packet{}
			if err := rtpPacket.Unmarshal(packet.Payload); err != nil {
				c.logger.Warn("Failed to unmarshal RTP packet %d: %v", i, err)
				continue
			}

			// Push packet to samplebuilder for reordering
			c.sampleBuilder.Push(rtpPacket)
		}

		// Log progress
		if i%2000 == 0 && i > 0 {
			c.logger.Info("Processed %d packets", i)
		}
	}

	if c.sampleBuilder != nil {
		c.sampleBuilder.Flush()
	}

	duration := time.Since(startTime)
	c.logger.Info("Finished feeding %d packets in %v", i, duration)

	// Allow some time for the recorder to finalize
	time.Sleep(2 * time.Second)

	return nil
}

func (c *RTPDump2WebMConverter) buildDefaultReleasePacketHandler() func(pkt *rtp.Packet) {
	return func(pkt *rtp.Packet) {
		if e := c.recorder.OnRTP(pkt); e != nil {
			c.logger.Warn("Failed to record RTP packet SeqNum: %d RtpTs: %d: %v", pkt.SequenceNumber, pkt.Timestamp, e)
		}
	}
}

func (c *RTPDump2WebMConverter) buildOpusReleasePacketHandler() func(pkt *rtp.Packet) {
	return func(pkt *rtp.Packet) {
		c.opusPacketDurationMsCorrected(pkt)

		pkt.SequenceNumber += c.inserted

		if false && c.lastPkt != nil {
			if pkt.SequenceNumber-c.lastPkt.SequenceNumber > 1 {
				c.logger.Info("Missing Packet Detected, Previous SeqNum: %d RtpTs: %d   - Last SeqNum: %d RtpTs: %d", c.lastPkt.SequenceNumber, c.lastPkt.Timestamp, pkt.SequenceNumber, pkt.Timestamp)
			}

			tsDiff := pkt.Timestamp - c.lastPkt.Timestamp // TODO handle rollover
			lastPktDuration := c.opusPacketDurationMsCorrected(c.lastPkt)
			rtpDuration := uint32(lastPktDuration * 48)

			if rtpDuration == 0 {
				rtpDuration = c.lastPktDuration
				c.logger.Info("LastPacket with no duration, Previous SeqNum: %d RtpTs: %d   - Last SeqNum: %d RtpTs: %d", c.lastPkt.SequenceNumber, c.lastPkt.Timestamp, pkt.SequenceNumber, pkt.Timestamp)
			} else {
				c.lastPktDuration = rtpDuration
			}

			if rtpDuration > 0 && tsDiff > rtpDuration {

				// Calculate how many packets we need to insert, taking care of packet losses
				var toAdd uint16
				if uint32(pkt.SequenceNumber-c.lastPkt.SequenceNumber)*rtpDuration != tsDiff { // TODO handle rollover
					toAdd = uint16(tsDiff/rtpDuration) - (pkt.SequenceNumber - c.lastPkt.SequenceNumber)
				}

				c.logger.Info("Gap detected, inserting %d packets tsDiff %d, Previous SeqNum: %d RtpTs: %d   - Last SeqNum: %d RtpTs: %d",
					toAdd, tsDiff, c.lastPkt.SequenceNumber, c.lastPkt.Timestamp, pkt.SequenceNumber, pkt.Timestamp)

				for i := 1; i <= int(toAdd); i++ {
					ins := c.lastPkt.Clone()
					ins.Payload = ins.Payload[:1] // Keeping only TOC byte
					ins.SequenceNumber += uint16(i)
					ins.Timestamp += uint32(i) * rtpDuration

					c.logger.Debug("Writing inserted Packet %v", ins)
					if e := c.recorder.OnRTP(ins); e != nil {
						c.logger.Warn("Failed to record inserted RTP packet SeqNum: %d RtpTs: %d: %v", ins.SequenceNumber, ins.Timestamp, e)
					}
				}

				c.inserted += toAdd
				pkt.SequenceNumber += toAdd
			}
		}

		c.lastPkt = pkt

		c.logger.Debug("Writing real Packet Last SeqNum: %d RtpTs: %d", pkt.SequenceNumber, pkt.Timestamp)
		if e := c.recorder.OnRTP(pkt); e != nil {
			c.logger.Warn("Failed to record RTP packet SeqNum: %d RtpTs: %d: %v", pkt.SequenceNumber, pkt.Timestamp, e)
		}
	}
}

func opusPacketDurationMs(packet []byte) int {
	if len(packet) < 1 {
		return 0
	}
	toc := packet[0]
	config := toc & 0x1F
	c := (toc >> 6) & 0x03

	frameDuration := 0
	switch {
	case config <= 3:
		frameDuration = 10 << (config & 0x03) // 10,20,40,60
	case config <= 7:
		frameDuration = 10 << (config & 0x03)
	case config <= 11:
		frameDuration = 10 << (config & 0x03)
	case config <= 13:
		frameDuration = 10 << (config & 0x01)
	case config <= 19:
		frameDuration = 25 / 10 // 2.5ms
	case config <= 23:
		frameDuration = 5
	case config <= 27:
		frameDuration = 10
	default:
		frameDuration = 20
	}

	var frameCount int
	switch c {
	case 0:
		frameCount = 1
	case 1, 2:
		frameCount = 2
	case 3:
		if len(packet) > 1 {
			frameCount = int(packet[1] & 0x3F)
		}
	}

	return frameDuration * frameCount
}

// opusPacketDurationMsCorrected implements the correct OPUS RFC 6716 specification
// with comprehensive logging to debug duration calculation issues
func (c *RTPDump2WebMConverter) opusPacketDurationMsCorrected(pkt *rtp.Packet) int {
	payload := pkt.Payload
	if len(payload) < 1 {
		c.logger.Warn("OPUS packet too short: %d bytes", len(payload))
		return 0
	}

	toc := payload[0]

	// Skip special OPUS packets (DTX, padding, etc.)
	// TOC=0xF8 is a special packet type that should be excluded entirely
	if toc == 0xF8 {
		c.logger.Info("OPUS: Skipping special packet TOC=0x%02X RTP=%d seq=%d", toc, pkt.Timestamp, pkt.SequenceNumber)
		return 0
	}

	// Also skip the specific 3-byte sequence if it exists
	if len(payload) >= 3 && payload[0] == 0xF8 && payload[1] == 0xFF && payload[2] == 0xFE {
		c.logger.Info("OPUS: Skipping special packet payload=[0x%02X 0x%02X 0x%02X] RTP=%d seq=%d", payload[0], payload[1], payload[2], pkt.Timestamp, pkt.SequenceNumber)
		return 0
	}

	config := toc & 0x1F
	frameCountCode := (toc >> 6) & 0x03

	// Calculate frame duration according to OPUS RFC 6716
	// The frame duration depends on the config value, not a simple formula
	var frameDuration int
	switch {
	case config <= 3:
		// SILK mode: 10, 20, 40, 60 ms
		frameDuration = 10 * (1 << config)
	case config <= 7:
		// SILK mode: 10, 20, 40, 60 ms
		frameDuration = 10 * (1 << (config & 0x03))
	case config <= 11:
		// SILK mode: 10, 20, 40, 60 ms
		frameDuration = 10 * (1 << (config & 0x03))
	case config <= 13:
		// SILK mode: 10, 20 ms
		frameDuration = 10 * (1 << (config & 0x01))
	case config <= 19:
		// CELT mode: 2.5, 5, 10, 20 ms
		frameDuration = 25 * (1 << (config & 0x03)) // 2.5ms * 10 for integer math
	case config <= 23:
		// CELT mode: 5, 10, 20, 40 ms
		frameDuration = 50 * (1 << (config & 0x03)) // 5ms * 10 for integer math
	case config <= 27:
		// CELT mode: 10, 20, 40, 80 ms
		frameDuration = 100 * (1 << (config & 0x03)) // 10ms * 10 for integer math
	default:
		// Default case
		frameDuration = 200 // 20ms * 10 for integer math
	}

	var frameCount int
	var secondByte uint8
	var secondByteBinary string

	switch frameCountCode {
	case 0:
		frameCount = 1
	case 1:
		frameCount = 2
	case 2:
		frameCount = 2
	case 3:
		if len(payload) > 1 {
			secondByte = payload[1]
			secondByteBinary = fmt.Sprintf("%08b", secondByte)
			// Frame count is in lower 6 bits, 0-indexed
			frameCount = int(secondByte&0x3F) + 1
		} else {
			c.logger.Warn("Frame count code 3 but no second byte available")
			frameCount = 1
		}
	default:
		c.logger.Warn("Invalid frame count code: %d", frameCountCode)
		frameCount = 1
	}

	totalDuration := frameDuration * frameCount

	// Convert back to actual milliseconds for logging
	frameDurationMs := frameDuration / 10
	totalDurationMs := totalDuration / 10

	if totalDurationMs != 20 {
		// Compact inline logging for debugging
		if frameCountCode == 3 && len(payload) > 1 {
			c.logger.Info("OPUS: TOC=0x%02X cfg=%d fcc=%d 2nd=0x%02X(%s) fc=%d fd=%dms td=%dms RTP=%d seq=%d",
				toc, config, frameCountCode, secondByte, secondByteBinary, frameCount, frameDurationMs, totalDurationMs, pkt.Timestamp, pkt.SequenceNumber)
		} else {
			c.logger.Info("OPUS: TOC=0x%02X cfg=%d fcc=%d fc=%d fd=%dms td=%dms RTP=%d seq=%d",
				toc, config, frameCountCode, frameCount, frameDurationMs, totalDurationMs, pkt.Timestamp, pkt.SequenceNumber)
		}

	}

	return totalDuration
}
