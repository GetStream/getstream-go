package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/GetStream/getstream-go/v3"
	"github.com/GetStream/getstream-go/v3/cmd/raw-recording-tools/webm"
)

type AudioMixer struct {
	logger *getstream.DefaultLogger
}

func NewAudioMixer(logger *getstream.DefaultLogger) *AudioMixer {
	return &AudioMixer{logger: logger}
}

// mixAllAudioTracks orchestrates the entire audio mixing workflow using existing extraction logic
func (p *AudioMixer) mixAllAudioTracks(globalArgs *GlobalArgs, mixAudioArgs *MixAudioArgs, metadata *RecordingMetadata, logger *getstream.DefaultLogger) error {
	// Step 1: Extract all matching audio tracks using existing extractTracks function
	logger.Info("Step 1/2: Extracting all matching audio tracks...")

	mediaFilter := "user"
	if mixAudioArgs.IncludeScreenShare {
		mediaFilter = "both"
	}

	err := extractTracks(globalArgs.WorkDir, globalArgs.Output, "", "", "", metadata, "audio", mediaFilter, true, true, logger)
	if err != nil {
		return fmt.Errorf("failed to extract audio tracks: %w", err)
	}

	fileOffsetMap := p.offset(metadata, logger)
	if len(fileOffsetMap) == 0 {
		return fmt.Errorf("no audio files were extracted - check your filter criteria")
	}

	logger.Info("Found %d extracted audio files to mix", len(fileOffsetMap))

	// Step 3: Mix all discovered audio files using existing webm.MixAudioFiles
	outputFile := filepath.Join(globalArgs.Output, "mixed_audio.webm")

	err = webm.MixAudioFiles(outputFile, fileOffsetMap, logger)
	if err != nil {
		return fmt.Errorf("failed to mix audio files: %w", err)
	}

	logger.Info("Successfully created mixed audio file: %s", outputFile)

	//// Clean up individual audio files (optional)
	//for _, audioFile := range audioFiles {
	//	if err := os.Remove(audioFile.FilePath); err != nil {
	//		logger.Warn("Failed to clean up temporary file %s: %v", audioFile.FilePath, err)
	//	}
	//}

	return nil
}

func (p *AudioMixer) offset(metadata *RecordingMetadata, logger *getstream.DefaultLogger) map[string]int64 {
	offsetMap := make(map[string]int64)

	sort.Slice(metadata.Tracks, func(i, j int) bool {
		return metadata.Tracks[i].Segments[0].metadata.FirstRtpUnixTimestamp < metadata.Tracks[j].Segments[0].metadata.FirstRtpUnixTimestamp
	})

	var firstTrack *TrackInfo
	for _, t := range metadata.Tracks {
		if t.TrackType == "audio" && !t.IsScreenshare {
			if firstTrack == nil {
				firstTrack = t
				offsetMap[t.ConcatenatedContainerPath] = 0
			} else {
				offset, err := calculateSyncOffsetFromFiles(firstTrack, t, logger)
				if err != nil {
					logger.Warn("Failed to calculate sync offset for audio tracks: %v", err)
					continue
				}
				offsetMap[t.ConcatenatedContainerPath] = offset
			}
		}
	}

	return offsetMap
}
