package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GetStream/getstream-go/v3"
)

type ExtractAudioArgs struct {
	UserID    string
	SessionID string
	TrackID   string
	FillGaps  bool
}

func runExtractAudio(args []string, globalArgs *GlobalArgs) {
	// Parse command-specific flags
	fs := flag.NewFlagSet("extract-audio", flag.ExitOnError)
	extractAudioArgs := &ExtractAudioArgs{}
	fs.StringVar(&extractAudioArgs.UserID, "userId", "", "Specify a userId (empty for all)")
	fs.StringVar(&extractAudioArgs.SessionID, "sessionId", "", "Specify a sessionId (empty for all)")
	fs.StringVar(&extractAudioArgs.TrackID, "trackId", "", "Specify a trackId (empty for all)")
	fs.BoolVar(&extractAudioArgs.FillGaps, "fill_gaps", false, "Fix DTX shrink audio, and fill with silence when track was muted")

	// Check for help flag before parsing
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			printExtractAudioUsage()
			return
		}
	}

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	// Validate input arguments against actual recording data
	metadata, err := validateInputArgs(globalArgs, extractAudioArgs.UserID, extractAudioArgs.SessionID, extractAudioArgs.TrackID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Validation error: %v\n", err)
		if globalArgs.InputFile != "" {
			fmt.Fprintf(os.Stderr, "\nTip: Use 'raw-tools --inputFile %s --output %s list-tracks --format users' to see available user IDs\n",
				globalArgs.InputFile, globalArgs.Output)
		}
		os.Exit(1)
	}

	// Setup logger
	logger := setupLogger(globalArgs.Verbose)

	logger.Info("Starting extract-audio command")

	fmt.Printf("Extract audio command with hierarchical filtering:\n")
	if globalArgs.InputFile != "" {
		fmt.Printf("  Input file: %s\n", globalArgs.InputFile)
	}
	if globalArgs.InputS3 != "" {
		fmt.Printf("  Input S3: %s\n", globalArgs.InputS3)
	}
	fmt.Printf("  Output directory: %s\n", globalArgs.Output)
	fmt.Printf("  User ID filter: %s\n", extractAudioArgs.UserID)

	if extractAudioArgs.TrackID != "" {
		fmt.Printf("  → Processing specific track '%s'\n", extractAudioArgs.TrackID)
	} else if extractAudioArgs.SessionID != "" {
		fmt.Printf("  → Processing all audio tracks for session '%s'\n", extractAudioArgs.SessionID)
	} else if extractAudioArgs.UserID != "" {
		fmt.Printf("  → Processing all audio tracks for user '%s'\n", extractAudioArgs.UserID)
	} else {
		fmt.Printf("  → Processing all audio tracks (no filters)\n")
	}
	fmt.Printf("  Fill gaps: %t\n", extractAudioArgs.FillGaps)

	// Implement extract audio functionality
	if err := extractAudioTracks(globalArgs, extractAudioArgs, metadata, logger); err != nil {
		logger.Error("Failed to extract audio: %v", err)
	}

	logger.Info("Extract audio command completed")
}

func printExtractAudioUsage() {
	fmt.Fprintf(os.Stderr, "Usage: raw-tools [global options] extract-audio [command options]\n\n")
	fmt.Fprintf(os.Stderr, "Generate playable audio files from raw recording tracks.\n")
	fmt.Fprintf(os.Stderr, "Supports formats: webm, mp3, and others.\n\n")
	fmt.Fprintf(os.Stderr, "Command Options (Hierarchical Filtering):\n")
	fmt.Fprintf(os.Stderr, "  --userId <id>          Specify a userId or * for all (default: *)\n")
	fmt.Fprintf(os.Stderr, "  --sessionId <id>       Specify a sessionId or * for all (default: *)\n")
	fmt.Fprintf(os.Stderr, "                         Ignored if --userId=*\n")
	fmt.Fprintf(os.Stderr, "  --trackId <id>         Specify a trackId or * for all (default: *)\n")
	fmt.Fprintf(os.Stderr, "                         Ignored if --userId=* or --sessionId=*\n")
	fmt.Fprintf(os.Stderr, "  --fill_gaps            Fix DTX shrink audio, fill with silence when muted\n\n")
	fmt.Fprintf(os.Stderr, "Hierarchical Filtering Logic:\n")
	fmt.Fprintf(os.Stderr, "  --userId=*             → Extract ALL users, sessions, tracks (sessionId/trackId ignored)\n")
	fmt.Fprintf(os.Stderr, "  --userId=user1 --sessionId=*  → Extract ALL sessions/tracks for user1 (trackId ignored)\n")
	fmt.Fprintf(os.Stderr, "  --userId=user1 --sessionId=session1 --trackId=*  → Extract ALL tracks for user1/session1\n")
	fmt.Fprintf(os.Stderr, "  --userId=user1 --sessionId=session1 --trackId=track1  → Extract specific track\n\n")
	fmt.Fprintf(os.Stderr, "Examples:\n")
	fmt.Fprintf(os.Stderr, "  # Extract audio for all users (sessionId/trackId ignored)\n")
	fmt.Fprintf(os.Stderr, "  raw-tools --inputFile recording.zip --output ./out extract-audio --userId '*'\n\n")
	fmt.Fprintf(os.Stderr, "  # Extract audio for specific user, all sessions (trackId ignored)\n")
	fmt.Fprintf(os.Stderr, "  raw-tools --inputFile recording.zip --output ./out extract-audio --userId user123 --sessionId '*'\n\n")
	fmt.Fprintf(os.Stderr, "  # Extract audio for specific user/session, all tracks\n")
	fmt.Fprintf(os.Stderr, "  raw-tools --inputFile recording.zip --output ./out extract-audio --userId user123 --sessionId session456 --trackId '*'\n\n")
	fmt.Fprintf(os.Stderr, "Global Options: Use 'raw-tools --help' to see global options.\n")
}

func extractAudioTracks(globalArgs *GlobalArgs, extractAudioArgs *ExtractAudioArgs, metadata *RecordingMetadata, logger *getstream.DefaultLogger) error {
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

	return extractTracks(workingDir, globalArgs.Output, extractAudioArgs.UserID, extractAudioArgs.SessionID, extractAudioArgs.TrackID, metadata, "audio", "both", extractAudioArgs.FillGaps, logger)
}
