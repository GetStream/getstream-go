package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/GetStream/getstream-go/v3"
	"github.com/GetStream/getstream-go/v3/cmd/raw-recording-tools/webm"
)

type ProcessAllArgs struct {
	UserID    string
	SessionID string
	TrackID   string
}

func runProcessAll(args []string, globalArgs *GlobalArgs) {
	// Parse command-specific flags
	fs := flag.NewFlagSet("process-all", flag.ExitOnError)
	processAllArgs := &ProcessAllArgs{}
	fs.StringVar(&processAllArgs.UserID, "userId", "", "Specify a userId (empty for all)")
	fs.StringVar(&processAllArgs.SessionID, "sessionId", "", "Specify a sessionId (empty for all)")
	fs.StringVar(&processAllArgs.TrackID, "trackId", "", "Specify a trackId (empty for all)")

	// Check for help flag before parsing
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			printProcessAllUsage()
			return
		}
	}

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	// Validate global arguments
	if err := validateGlobalArgs(globalArgs, "process-all"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		printProcessAllUsage()
		os.Exit(1)
	}

	// Validate input arguments against actual recording data
	if err := validateInputArgs(globalArgs, processAllArgs.UserID, processAllArgs.SessionID, processAllArgs.TrackID); err != nil {
		fmt.Fprintf(os.Stderr, "Validation error: %v\n", err)
		if globalArgs.InputFile != "" {
			fmt.Fprintf(os.Stderr, "\nTip: Use 'raw-tools --inputFile %s --output %s list-tracks --format users' to see available user IDs\n",
				globalArgs.InputFile, globalArgs.Output)
		}
		os.Exit(1)
	}

	// Set up logger
	logger := setupLogger(globalArgs.Verbose)
	logger.Info("Starting process-all command")

	// Display hierarchy information for user clarity
	fmt.Printf("Process-all command (audio + video + mux) with hierarchical filtering:\n")
	fmt.Printf("  Input file: %s\n", globalArgs.InputFile)
	fmt.Printf("  Output directory: %s\n", globalArgs.Output)
	fmt.Printf("  User ID filter: %s\n", processAllArgs.UserID)
	fmt.Printf("  Session ID filter: %s\n", processAllArgs.SessionID)
	fmt.Printf("  Track ID filter: %s\n", processAllArgs.TrackID)
	fmt.Printf("  Gap filling: always enabled\n")

	if processAllArgs.UserID == "*" {
		fmt.Printf("  → Processing ALL users (sessionId/trackId ignored)\n")
	} else if processAllArgs.SessionID == "*" {
		fmt.Printf("  → Processing ALL sessions for user '%s' (trackId ignored)\n", processAllArgs.UserID)
	} else if processAllArgs.TrackID == "*" {
		fmt.Printf("  → Processing ALL tracks for user '%s', session '%s'\n", processAllArgs.UserID, processAllArgs.SessionID)
	} else {
		fmt.Printf("  → Processing specific track for user '%s', session '%s', track '%s'\n", processAllArgs.UserID, processAllArgs.SessionID, processAllArgs.TrackID)
	}

	// Process all tracks and mux them
	if err := processAllTracks(globalArgs, processAllArgs, logger); err != nil {
		logger.Error("Failed to process and mux tracks: %v", err)
		os.Exit(1)
	}

	logger.Info("Process-all command completed successfully")
}

func printProcessAllUsage() {
	fmt.Printf("Usage: process-all [OPTIONS]\n")
	fmt.Printf("\nProcess audio, video, and mux them into combined files (all-in-one workflow)\n")
	fmt.Printf("Outputs 3 files per session: audio WebM, video WebM, and muxed WebM\n")
	fmt.Printf("Gap filling is always enabled for seamless playback.\n")
	fmt.Printf("\nOptions:\n")
	fmt.Printf("  --userId STRING    Specify a userId or * for all (default: \"*\")\n")
	fmt.Printf("  --sessionId STRING Specify a sessionId or * for all (default: \"*\")\n")
	fmt.Printf("  --trackId STRING   Specify a trackId or * for all (default: \"*\")\n")
	fmt.Printf("\nOutput files per session:\n")
	fmt.Printf("  audio_{userId}_{sessionId}_{trackId}.webm    - Audio-only file\n")
	fmt.Printf("  video_{userId}_{sessionId}_{trackId}.webm    - Video-only file\n")
	fmt.Printf("  muxed_{userId}_{sessionId}_{trackId}.webm    - Combined audio+video file\n")
}

func processAllTracks(globalArgs *GlobalArgs, processAllArgs *ProcessAllArgs, logger *getstream.DefaultLogger) error {
	// Step 1: Extract audio tracks with gap filling
	logger.Info("Step 1/3: Extracting audio tracks with gap filling...")
	err := extractTracks(globalArgs, processAllArgs.UserID, processAllArgs.SessionID, processAllArgs.TrackID, "audio", "both", true, logger)
	if err != nil {
		return fmt.Errorf("failed to extract audio tracks: %w", err)
	}

	// Step 2: Extract video tracks with gap filling
	logger.Info("Step 2/3: Extracting video tracks with gap filling...")
	err = extractTracks(globalArgs, processAllArgs.UserID, processAllArgs.SessionID, processAllArgs.TrackID, "video", "both", true, logger)
	if err != nil {
		return fmt.Errorf("failed to extract video tracks: %w", err)
	}

	// Step 3: Mux audio and video files (keeping originals)
	logger.Info("Step 3/3: Muxing audio and video tracks...")
	err = muxAudioVideoTracksKeepOriginals(globalArgs, processAllArgs, logger)
	if err != nil {
		return fmt.Errorf("failed to mux audio and video tracks: %w", err)
	}

	// Report final output
	audioFiles, _ := filepath.Glob(filepath.Join(globalArgs.Output, "audio_*.webm"))
	videoFiles, _ := filepath.Glob(filepath.Join(globalArgs.Output, "video_*.webm"))
	muxedFiles, _ := filepath.Glob(filepath.Join(globalArgs.Output, "muxed_*.webm"))

	logger.Info("Process-all completed successfully:")
	logger.Info("  - %d audio files", len(audioFiles))
	logger.Info("  - %d video files", len(videoFiles))
	logger.Info("  - %d muxed files", len(muxedFiles))

	fmt.Printf("\n✅ Generated files in %s:\n", globalArgs.Output)
	for _, file := range audioFiles {
		fmt.Printf("  🎵 %s\n", filepath.Base(file))
	}
	for _, file := range videoFiles {
		fmt.Printf("  🎬 %s\n", filepath.Base(file))
	}
	for _, file := range muxedFiles {
		fmt.Printf("  🎞️  %s\n", filepath.Base(file))
	}

	return nil
}

// muxAudioVideoTracksKeepOriginals is like muxAudioVideoTracks but keeps the original audio/video files
func muxAudioVideoTracksKeepOriginals(globalArgs *GlobalArgs, processAllArgs *ProcessAllArgs, logger *getstream.DefaultLogger) error {
	// Find the generated audio and video WebM files
	audioFiles, err := filepath.Glob(filepath.Join(globalArgs.Output, "audio_*.webm"))
	if err != nil {
		return fmt.Errorf("failed to find audio files: %w", err)
	}
	if len(audioFiles) == 0 {
		return fmt.Errorf("no audio files generated")
	}

	videoFiles, err := filepath.Glob(filepath.Join(globalArgs.Output, "video_*.webm"))
	if err != nil {
		return fmt.Errorf("failed to find video files: %w", err)
	}
	if len(videoFiles) == 0 {
		return fmt.Errorf("no video files generated")
	}

	logger.Info("Found %d audio files and %d video files to mux", len(audioFiles), len(videoFiles))

	// Mux each audio/video pair
	for i, audioFile := range audioFiles {
		if i >= len(videoFiles) {
			logger.Warn("No matching video file for audio file %s", audioFile)
			continue
		}
		videoFile := videoFiles[i]

		// Calculate sync offset using segment timing information
		offset, err := calculateSyncOffsetFromFiles(globalArgs.InputFile, audioFile, videoFile, logger)
		if err != nil {
			logger.Warn("Failed to calculate sync offset, using 0: %v", err)
			offset = 0
		}

		// Generate output filename
		outputFile := generateMuxedFilename(audioFile, videoFile, globalArgs.Output)

		// Mux the audio and video files
		logger.Info("Muxing %s + %s → %s (offset: %dms)",
			filepath.Base(audioFile), filepath.Base(videoFile), filepath.Base(outputFile), offset)

		err = webm.MuxFiles(outputFile, audioFile, videoFile, float64(offset), logger)
		if err != nil {
			logger.Error("Failed to mux %s + %s: %v", audioFile, videoFile, err)
			continue
		}

		logger.Info("Successfully created muxed file: %s", outputFile)
		// NOTE: Unlike muxAudioVideoTracks, we DON'T clean up the individual files here
	}

	return nil
}
