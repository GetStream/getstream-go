package rawrecorder

import "time"

type SessionTimingMetadata struct {
	CallID        string    `json:"call_id"`
	CallSessionID string    `json:"call_session_id"`
	CallStartTime time.Time `json:"call_start_time"`
	CallEndTime   time.Time `json:"call_end_time,omitempty"`
	ParticipantID string    `json:"participant_id"`
	UserSessionID string    `json:"user_session_id"`
	Segments      struct {
		Audio []SegmentMetadata `json:"audio"`
		Video []SegmentMetadata `json:"video"`
	} `json:"segments"`
}

type SegmentMetadata struct {
	// Global information
	BaseFilename   string `json:"base_filename"`
	StartTimestamp int64  `json:"start_timestamp"`
	EndTimestamp   int64  `json:"end_timestamp,omitempty"`
	StartOffsetMs  int64  `json:"start_offset_ms"`
	EndOffsetMs    int64  `json:"end_offset_ms,omitempty"`

	// Packet timing information
	FirstRtpRtpTimestamp  uint32 `json:"first_rtp_rtp_timestamp"`
	FirstRtpUnixTimestamp int64  `json:"first_rtp_unix_timestamp"`
	LastRtpRtpTimestamp   uint32 `json:"last_rtp_rtp_timestamp,omitempty"`
	LastRtpUnixTimestamp  int64  `json:"last_rtp_unix_timestamp,omitempty"`
	FirstRtcpRtpTimestamp uint32 `json:"first_rtcp_rtp_timestamp,omitempty"`
	FirstRtcpNtpTimestamp int64  `json:"first_rtcp_ntp_timestamp,omitempty"`
	LastRtcpRtpTimestamp  uint32 `json:"last_rtcp_rtp_timestamp,omitempty"`
	LastRtcpNtpTimestamp  int64  `json:"last_rtcp_ntp_timestamp,omitempty"`

	// Segment duration information
	RtpDurationMs    int64   `json:"rtp_duration_ms,omitempty"`
	UnixDurationMs   int64   `json:"unix_duration_ms,omitempty"`
	DriftMs          int64   `json:"drift_ms,omitempty"`
	DriftRatePercent float64 `json:"drift_rate_percent,omitempty"`

	// Track information
	SSRC      uint32 `json:"ssrc"`
	Codec     string `json:"codec"`
	TrackID   string `json:"track_id"`
	TrackType string `json:"track_type"`
}
