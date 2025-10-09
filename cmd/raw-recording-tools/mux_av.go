package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GetStream/getstream-go/v3"
	"github.com/GetStream/getstream-go/v3/cmd/raw-recording-tools/webm"
)

type MuxAVArgs struct {
	UserID    string
	SessionID string
	TrackID   string
	Media     string // "user", "display", or "both" (default)
}

func runMuxAV(args []string, globalArgs *GlobalArgs) {
	// Parse command-specific flags
	fs := flag.NewFlagSet("mux-av", flag.ExitOnError)
	muxAVArgs := &MuxAVArgs{}
	fs.StringVar(&muxAVArgs.UserID, "userId", "", "Specify a userId (empty for all)")
	fs.StringVar(&muxAVArgs.SessionID, "sessionId", "", "Specify a sessionId (empty for all)")
	fs.StringVar(&muxAVArgs.TrackID, "trackId", "", "Specify a trackId (empty for all)")
	fs.StringVar(&muxAVArgs.Media, "media", "both", "Filter by media type: 'user', 'display', or 'both'")

	// Check for help flag before parsing
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			printMuxAVUsage()
			return
		}
	}

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	// Validate input arguments against actual recording data
	metadata, err := validateInputArgs(globalArgs, muxAVArgs.UserID, muxAVArgs.SessionID, muxAVArgs.TrackID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Validation error: %v\n", err)
		if globalArgs.InputFile != "" {
			fmt.Fprintf(os.Stderr, "\nTip: Use 'raw-tools --inputFile %s --output %s list-tracks --format users' to see available user IDs\n",
				globalArgs.InputFile, globalArgs.Output)
		}
		os.Exit(1)
	}

	// Set up logger
	logger := setupLogger(globalArgs.Verbose)
	logger.Info("Starting mux-av command")

	// Display hierarchy information for user clarity
	fmt.Printf("Mux audio and video command with hierarchical filtering:\n")
	fmt.Printf("  Input file: %s\n", globalArgs.InputFile)
	fmt.Printf("  Output directory: %s\n", globalArgs.Output)
	fmt.Printf("  User ID filter: %s\n", muxAVArgs.UserID)
	fmt.Printf("  Session ID filter: %s\n", muxAVArgs.SessionID)
	fmt.Printf("  Track ID filter: %s\n", muxAVArgs.TrackID)
	fmt.Printf("  Media filter: %s\n", muxAVArgs.Media)

	if muxAVArgs.TrackID != "" {
		fmt.Printf("  → Processing specific track '%s'\n", muxAVArgs.TrackID)
	} else if muxAVArgs.SessionID != "" {
		fmt.Printf("  → Processing all tracks for session '%s'\n", muxAVArgs.SessionID)
	} else if muxAVArgs.UserID != "" {
		fmt.Printf("  → Processing all tracks for user '%s'\n", muxAVArgs.UserID)
	} else {
		fmt.Printf("  → Processing all tracks (no filters)\n")
	}

	// Extract and mux audio/video tracks
	if err := muxAudioVideoTracks(globalArgs, muxAVArgs, metadata, logger); err != nil {
		logger.Error("Failed to mux audio/video tracks: %v", err)
		os.Exit(1)
	}

	logger.Info("Mux audio and video command completed successfully")
}

func printMuxAVUsage() {
	fmt.Printf("Usage: mux-av [OPTIONS]\n")
	fmt.Printf("\nMux audio and video tracks into a single file\n")
	fmt.Printf("\nOptions:\n")
	fmt.Printf("  --userId STRING    Specify a userId or * for all (default: \"*\")\n")
	fmt.Printf("  --sessionId STRING Specify a sessionId or * for all (default: \"*\")\n")
	fmt.Printf("  --trackId STRING   Specify a trackId or * for all (default: \"*\")\n")
	fmt.Printf("  --media STRING     Filter by media type: 'user', 'display', or 'both' (default: \"both\")\n")
	fmt.Printf("\nMedia Filtering:\n")
	fmt.Printf("  --media user     Only mux user camera audio/video pairs\n")
	fmt.Printf("  --media display  Only mux display sharing audio/video pairs\n")
	fmt.Printf("  --media both     Mux both types, but ensure consistent pairing (default)\n")
}

func muxAudioVideoTracks(globalArgs *GlobalArgs, muxAVArgs *MuxAVArgs, metadata *RecordingMetadata, logger *getstream.DefaultLogger) error {
	// Extract to temp directory if needed (unified approach)
	workingDir, cleanup, err := extractToTempDir(globalArgs.InputFile, logger)
	if err != nil {
		return fmt.Errorf("failed to prepare working directory: %w", err)
	}
	defer cleanup()

	// Create output directory if it doesn't exist
	if e := os.MkdirAll(globalArgs.Output, 0755); e != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create a temporary directory for intermediate files
	tempDir, err := os.MkdirTemp("", "mux-av-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Extract audio tracks with gap filling enabled
	logger.Info("Extracting audio tracks with gap filling...")
	err = extractTracks(workingDir, globalArgs.Output, muxAVArgs.UserID, muxAVArgs.SessionID, muxAVArgs.TrackID, metadata, "audio", muxAVArgs.Media, true, logger)
	if err != nil {
		return fmt.Errorf("failed to extract audio tracks: %w", err)
	}

	// Extract video tracks with gap filling enabled
	logger.Info("Extracting video tracks with gap filling...")
	err = extractTracks(workingDir, globalArgs.Output, muxAVArgs.UserID, muxAVArgs.SessionID, muxAVArgs.TrackID, metadata, "video", muxAVArgs.Media, true, logger)
	if err != nil {
		return fmt.Errorf("failed to extract video tracks: %w", err)
	}

	// Find the generated audio and video WebM files
	audioFiles, err := filepath.Glob(filepath.Join(globalArgs.Output, "audio_*.webm"))
	if err != nil {
		return fmt.Errorf("failed to find audio files: %w", err)
	}
	if len(audioFiles) == 0 {
		return fmt.Errorf("no audio files generated")
	}

	webmVideoFiles, err := filepath.Glob(filepath.Join(globalArgs.Output, "video_*.webm"))
	mp4VideoFiles, err := filepath.Glob(filepath.Join(globalArgs.Output, "video_*.mp4"))

	videoFiles := append(webmVideoFiles, mp4VideoFiles...)

	if err != nil {
		return fmt.Errorf("failed to find video files: %w", err)
	}
	if len(videoFiles) == 0 {
		return fmt.Errorf("no video files generated")
	}

	logger.Info("Found %d audio files and %d video files to mux", len(audioFiles), len(videoFiles))

	// Group files by media type for proper pairing
	audioGroups, videoGroups, err := groupFilesByMediaType(globalArgs.InputFile, audioFiles, videoFiles, muxAVArgs.Media, metadata, logger)
	if err != nil {
		return fmt.Errorf("failed to group files by media type: %w", err)
	}

	// Mux user tracks
	if userAudio, userVideo := audioGroups["user"], videoGroups["user"]; len(userAudio) > 0 && len(userVideo) > 0 {
		logger.Info("Muxing %d user audio/video pairs", len(userAudio))
		err = muxTrackPairs(globalArgs.InputFile, userAudio, userVideo, globalArgs.Output, "user", metadata, logger)
		if err != nil {
			logger.Error("Failed to mux user tracks: %v", err)
		}
	}

	// Mux display tracks
	if displayAudio, displayVideo := audioGroups["display"], videoGroups["display"]; len(displayAudio) > 0 && len(displayVideo) > 0 {
		logger.Info("Muxing %d display audio/video pairs", len(displayAudio))
		err = muxTrackPairs(globalArgs.InputFile, displayAudio, displayVideo, globalArgs.Output, "display", metadata, logger)
		if err != nil {
			logger.Error("Failed to mux display tracks: %v", err)
		}
	}

	return nil
}

// calculateSyncOffsetFromFiles calculates sync offset between audio and video files using metadata
func calculateSyncOffsetFromFiles(inputPath, audioFile, videoFile string, metadata *RecordingMetadata, logger *getstream.DefaultLogger) (int64, error) {
	// Extract track IDs from filenames
	audioTrackID := extractTrackIDFromFilename(audioFile)
	videoTrackID := extractTrackIDFromFilename(videoFile)

	if audioTrackID == "" || videoTrackID == "" {
		return 0, fmt.Errorf("could not extract track IDs from filenames")
	}

	// Find the audio and video tracks
	var audioTrack, videoTrack *TrackInfo
	for _, track := range metadata.Tracks {
		if track.TrackID == audioTrackID && track.TrackType == "audio" {
			audioTrack = track
		}
		if track.TrackID == videoTrackID && track.TrackType == "video" {
			videoTrack = track
		}
	}

	if audioTrack == nil || videoTrack == nil {
		return 0, fmt.Errorf("could not find matching tracks in metadata")
	}

	// Calculate offset: positive means video starts before audio
	audioTs := FirstPacketNtpTimestamp(audioTrack.Segments[0])
	videoTs := FirstPacketNtpTimestamp(videoTrack.Segments[0])
	offset := audioTs - videoTs

	logger.Info(fmt.Sprintf("Calculated sync offset: audio_start=%v, audio_ts=%v, video_start=%v, video_ts=%v, offset=%d",
		audioTrack.Segments[0].FirstRtpUnixTimestamp, audioTs, videoTrack.Segments[0].FirstRtpUnixTimestamp, videoTs, offset))

	return offset, nil
}

// extractTrackIDFromFilename extracts track ID from generated filename
func extractTrackIDFromFilename(filename string) string {
	// Filename format: {type}_{userId}_{sessionId}_{trackId}.suffix
	base := filepath.Base(filename)
	split := strings.Split(base, ".")
	parts := strings.Split(split[0], "_")
	if len(parts) >= 4 {
		return parts[3] // trackId is the 4th part
	}
	return ""
}

// generateMuxedFilename creates output filename for muxed file
func generateMuxedFilename(audioFile, videoFile, outputDir string) string {
	// Extract common parts from audio filename
	videoBase := filepath.Base(videoFile)
	split := strings.Split(videoBase, ".")

	// Replace "audio_" with "muxed_" to create output name
	muxedName := strings.Replace(split[0], "audio_", "muxed_", 1) + "." + split[1]

	return filepath.Join(outputDir, muxedName)
}

// groupFilesByMediaType groups audio and video files by media type (user vs display)
func groupFilesByMediaType(inputPath string, audioFiles, videoFiles []string, mediaFilter string, metadata *RecordingMetadata, logger *getstream.DefaultLogger) (map[string][]string, map[string][]string, error) {
	// Create track ID to screenshare type mapping
	trackScreenshareMap := make(map[string]bool)
	for _, track := range metadata.Tracks {
		trackScreenshareMap[track.TrackID] = track.IsScreenshare
	}

	// Group files by media type
	audioGroups := map[string][]string{"user": {}, "display": {}}
	videoGroups := map[string][]string{"user": {}, "display": {}}

	// Process audio files
	for _, audioFile := range audioFiles {
		trackID := extractTrackIDFromFilename(audioFile)
		if trackID == "" {
			logger.Warn("Could not extract track ID from audio file: %s", audioFile)
			continue
		}

		isScreenshare, exists := trackScreenshareMap[trackID]
		if !exists {
			logger.Warn("Track ID %s not found in metadata for audio file: %s", trackID, audioFile)
			continue
		}

		// Apply media filter
		if mediaFilter == "user" && isScreenshare {
			continue // Skip display tracks when only user requested
		}
		if mediaFilter == "display" && !isScreenshare {
			continue // Skip user tracks when only display requested
		}

		if isScreenshare {
			audioGroups["display"] = append(audioGroups["display"], audioFile)
		} else {
			audioGroups["user"] = append(audioGroups["user"], audioFile)
		}
	}

	// Process video files
	for _, videoFile := range videoFiles {
		trackID := extractTrackIDFromFilename(videoFile)
		if trackID == "" {
			logger.Warn("Could not extract track ID from video file: %s", videoFile)
			continue
		}

		isScreenshare, exists := trackScreenshareMap[trackID]
		if !exists {
			logger.Warn("Track ID %s not found in metadata for video file: %s", trackID, videoFile)
			continue
		}

		// Apply media filter
		if mediaFilter == "user" && isScreenshare {
			continue // Skip display tracks when only user requested
		}
		if mediaFilter == "display" && !isScreenshare {
			continue // Skip user tracks when only display requested
		}

		if isScreenshare {
			videoGroups["display"] = append(videoGroups["display"], videoFile)
		} else {
			videoGroups["user"] = append(videoGroups["user"], videoFile)
		}
	}

	logger.Info("Grouped files: user audio=%d, user video=%d, display audio=%d, display video=%d",
		len(audioGroups["user"]), len(videoGroups["user"]),
		len(audioGroups["display"]), len(videoGroups["display"]))

	return audioGroups, videoGroups, nil
}

// muxTrackPairs muxes audio/video pairs of the same media type
func muxTrackPairs(inputPath string, audioFiles, videoFiles []string, outputDir, mediaTypeName string, metadata *RecordingMetadata, logger *getstream.DefaultLogger) error {
	minLen := len(audioFiles)
	if len(videoFiles) < minLen {
		minLen = len(videoFiles)
	}

	if minLen == 0 {
		logger.Warn("No %s audio/video pairs to mux", mediaTypeName)
		return nil
	}

	for i := 0; i < minLen; i++ {
		audioFile := audioFiles[i]
		videoFile := videoFiles[i]

		// Calculate sync offset using segment timing information
		offset, err := calculateSyncOffsetFromFiles(inputPath, audioFile, videoFile, metadata, logger)
		if err != nil {
			logger.Warn("Failed to calculate sync offset, using 0: %v", err)
			offset = 0
		}

		// Generate output filename with media type indicator
		outputFile := generateMediaAwareMuxedFilename(audioFile, videoFile, outputDir, mediaTypeName)

		// Mux the audio and video files
		logger.Info("Muxing %s %s + %s → %s (offset: %dms)",
			mediaTypeName, filepath.Base(audioFile), filepath.Base(videoFile), filepath.Base(outputFile), offset)

		err = webm.MuxFiles(outputFile, audioFile, videoFile, float64(offset), logger)
		if err != nil {
			logger.Error("Failed to mux %s + %s: %v", audioFile, videoFile, err)
			continue
		}

		logger.Info("Successfully created %s muxed file: %s", mediaTypeName, outputFile)

		// Clean up individual track files to avoid clutter
		//os.Remove(audioFile)
		//os.Remove(videoFile)
	}

	if len(audioFiles) != len(videoFiles) {
		logger.Warn("Mismatched %s track counts: %d audio, %d video", mediaTypeName, len(audioFiles), len(videoFiles))
	}

	return nil
}

// generateMediaAwareMuxedFilename creates output filename that indicates media type
func generateMediaAwareMuxedFilename(audioFile, videoFile, outputDir, mediaTypeName string) string {
	audioBase := filepath.Base(audioFile)
	audioBase = strings.TrimSuffix(audioBase, ".webm")

	// Extract common parts from audio filename
	videoBase := filepath.Base(videoFile)
	split := strings.Split(videoBase, ".")

	// Replace "audio_" with "muxed_{mediaType}_" to create output name
	var muxedName string
	if mediaTypeName == "display" {
		muxedName = strings.Replace(audioBase, "audio_", "muxed_display_", 1) + "." + split[1]
	} else {
		muxedName = strings.Replace(audioBase, "audio_", "muxed_", 1) + "." + split[1]
	}

	return filepath.Join(outputDir, muxedName)
}
