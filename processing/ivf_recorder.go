package processing

import (
	"os"
	"time"

	"github.com/pion/rtp"
	"github.com/pion/webrtc/v4/pkg/media/ivfwriter"
)

type IvfDumpRecorder struct {
	file      *os.File
	startTime time.Time
	writer    *ivfwriter.IVFWriter
}

func NewIvfDumpRecorder(outputPath string, mimeType string) (*IvfDumpRecorder, error) {
	writer, _ := ivfwriter.New(outputPath, ivfwriter.WithCodec(mimeType))

	recorder := &IvfDumpRecorder{
		startTime: time.Now(),
		writer:    writer,
	}

	return recorder, nil
}

func (r *IvfDumpRecorder) OnRTP(packet *rtp.Packet) error {
	err := r.writer.WriteRTP(packet)
	return err
}

func (r *IvfDumpRecorder) PushRtpBuf(buf []byte) error {
	rtpPacket := &rtp.Packet{}
	if err := rtpPacket.Unmarshal(buf); err != nil {
		return err
	}

	return r.OnRTP(rtpPacket)
}

func (r *IvfDumpRecorder) Close() error {
	return r.file.Close()
}
