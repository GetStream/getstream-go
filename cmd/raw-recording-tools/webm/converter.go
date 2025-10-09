package webm

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/GetStream/getstream-go/v3"
	"github.com/GetStream/getstream-go/v3/cmd/raw-recording-tools/rawsdputil"
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

	lastPkt  *rtp.Packet
	inserted uint16
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

func ConvertDirectory(directory string, logger *getstream.DefaultLogger) error {
	var rtpdumpFiles []string

	// Walk through directory to find .rtpdump files
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".rtpdump") {
			rtpdumpFiles = append(rtpdumpFiles, path)
		}

		return nil
	})
	if err != nil {
		return err
	}

	for _, rtpdumpFile := range rtpdumpFiles {
		c := newRTPDump2WebMConverter(logger)
		if err := c.ConvertFile(rtpdumpFile); err != nil {
			c.logger.Error("Failed to convert %s: %v", rtpdumpFile, err)
			continue
		}
	}

	return nil
}

func (c *RTPDump2WebMConverter) ConvertFile(inputFile string) error {
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

	sdpContent, _ := rawsdputil.ReadSDP(strings.Replace(inputFile, ".rtpdump", ".sdp", 1))
	mType, _ := rawsdputil.MimeType(sdpContent)

	var recorder WebmRecorder
	switch mType {
	case webrtc.MimeTypeAV1, webrtc.MimeTypeVP9:
		recorder, err = NewCursorGstreamerWebmRecorder(strings.Replace(inputFile, ".rtpdump", ".webm", 1), sdpContent, c.logger)
	case webrtc.MimeTypeH264:
		recorder, err = NewCursorWebmRecorder(strings.Replace(inputFile, ".rtpdump", ".mp4", 1), sdpContent, c.logger)
	default:
		recorder, err = NewCursorWebmRecorder(strings.Replace(inputFile, ".rtpdump", ".webm", 1), sdpContent, c.logger)
	}
	if err != nil {
		return fmt.Errorf("failed to create WebM recorder: %w", err)
	}
	defer recorder.Close()

	c.recorder = recorder

	options := samplebuilder.WithPacketReleaseHandler(func(pkt *rtp.Packet) {
		pkt.SequenceNumber += c.inserted

		if c.lastPkt != nil {
			if pkt.SequenceNumber-c.lastPkt.SequenceNumber > 1 {
				c.logger.Info("Missing Packet Detected, Previous SeqNum: %d RtpTs: %d   - Last SeqNum: %d RtpTs: %d", c.lastPkt.SequenceNumber, c.lastPkt.Timestamp, pkt.SequenceNumber, pkt.Timestamp)
			}

			if mType == webrtc.MimeTypeOpus {
				tsDiff := pkt.Timestamp - c.lastPkt.Timestamp // TODO handle rollover
				lastPktDuration := opusPacketDurationMs(c.lastPkt.Payload)
				rtpDuration := uint32(lastPktDuration * 48)
				if tsDiff > rtpDuration {

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
						e := c.recorder.OnRTP(ins)
						if e != nil {
							c.logger.Warn("Failed to record RTP packet %v: %v", pkt, err)
						}

						// Need to compute new packet
					}
					c.inserted += toAdd
					pkt.SequenceNumber += toAdd
					//				c.logger.Debugf("Inserted %d packets Previous inserting %s", toAdd, c.inserted)
				}
			}
		}

		c.lastPkt = pkt

		c.logger.Debug("Writing real Packet Last SeqNum: %d RtpTs: %d", pkt.SequenceNumber, pkt.Timestamp)
		e := c.recorder.OnRTP(pkt)
		if e != nil {
			c.logger.Warn("Failed to record RTP packet %v: %v", pkt, err)
		}
	})

	// Initialize samplebuilder based on codec type
	var sampleBuilder *samplebuilder.SampleBuilder
	switch mType {
	case webrtc.MimeTypeOpus:
		sampleBuilder = samplebuilder.New(audioMaxLate, &codecs.OpusPacket{}, 48000, options)
	case webrtc.MimeTypeVP8:
		sampleBuilder = samplebuilder.New(videoMaxLate, &codecs.VP8Packet{}, 90000, options)
	case webrtc.MimeTypeVP9:
		sampleBuilder = samplebuilder.New(videoMaxLate, &codecs.VP9Packet{}, 90000, options)
	case webrtc.MimeTypeH264:
		sampleBuilder = samplebuilder.New(videoMaxLate, &codecs.H264Packet{}, 90000, options)
	case webrtc.MimeTypeAV1:
		sampleBuilder = samplebuilder.New(videoMaxLate, &codecs.AV1Depacketizer{}, 90000, options)
	default:
		return fmt.Errorf("unsupported codec type: %s", mType)
	}
	c.sampleBuilder = sampleBuilder

	time.Sleep(1 * time.Second)

	// Convert and feed RTP packets
	return c.feedPackets(reader)
}

func (c *RTPDump2WebMConverter) feedPackets(reader *rtpdump.Reader) error {
	startTime := time.Now()

	for i := 0; ; i++ {
		packet, err := reader.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return err
			}
		}

		if packet.IsRTCP {
			continue
		}

		// Unmarshal the RTP packet from the raw payload

		if c.sampleBuilder == nil {
			_ = c.recorder.PushRtpBuf(packet.Payload)
		} else
		// Unmarshal the RTP packet from the raw payload
		{
			rtpPacket := &rtp.Packet{}
			if err := rtpPacket.Unmarshal(packet.Payload); err != nil {
				c.logger.Warn("Failed to unmarshal RTP packet %d: %v", i, err)
				continue
			}

			c.sampleBuilder.Push(rtpPacket)
		}
		// Push packet to samplebuilder for reordering

		// Log progress
		if i%100 == 0 && i > 0 {
			c.logger.Info("Processed %d packets", i)
		}
	}

	if c.sampleBuilder != nil {
		c.sampleBuilder.Flush()
	}

	duration := time.Since(startTime)
	c.logger.Info("Finished feeding packets in %v", duration)

	// Allow some time for the recorder to finalize
	time.Sleep(2 * time.Second)

	return nil
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
