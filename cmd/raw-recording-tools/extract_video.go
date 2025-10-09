package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GetStream/getstream-go/v3"
)

type ExtractVideoArgs struct {
	UserID    string
	SessionID string
	TrackID   string
	FillGaps  bool
}

func runExtractVideo(args []string, globalArgs *GlobalArgs) {
	printHelpIfAsked(args, printExtractVideoUsage)

	// Parse command-specific flags
	fs := flag.NewFlagSet("extract-video", flag.ExitOnError)
	extractVideoArgs := &ExtractVideoArgs{}
	fs.StringVar(&extractVideoArgs.UserID, "userId", "", "Specify a userId (empty for all)")
	fs.StringVar(&extractVideoArgs.SessionID, "sessionId", "", "Specify a sessionId (empty for all)")
	fs.StringVar(&extractVideoArgs.TrackID, "trackId", "", "Specify a trackId (empty for all)")
	fs.BoolVar(&extractVideoArgs.FillGaps, "fill_gaps", false, "Fill with black frame when track was muted")

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	// Validate input arguments against actual recording data
	metadata, err := validateInputArgs(globalArgs, extractVideoArgs.UserID, extractVideoArgs.SessionID, extractVideoArgs.TrackID)
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

	logger.Info("Starting extract-video command")

	fmt.Printf("Extract video command with hierarchical filtering:\n")
	if globalArgs.InputFile != "" {
		fmt.Printf("  Input file: %s\n", globalArgs.InputFile)
	}
	if globalArgs.InputS3 != "" {
		fmt.Printf("  Input S3: %s\n", globalArgs.InputS3)
	}
	fmt.Printf("  Output directory: %s\n", globalArgs.Output)
	fmt.Printf("  User ID filter: %s\n", extractVideoArgs.UserID)

	if extractVideoArgs.TrackID != "" {
		fmt.Printf("  → Processing specific track '%s'\n", extractVideoArgs.TrackID)
	} else if extractVideoArgs.SessionID != "" {
		fmt.Printf("  → Processing all video tracks for session '%s'\n", extractVideoArgs.SessionID)
	} else if extractVideoArgs.UserID != "" {
		fmt.Printf("  → Processing all video tracks for user '%s'\n", extractVideoArgs.UserID)
	} else {
		fmt.Printf("  → Processing all video tracks (no filters)\n")
	}
	fmt.Printf("  Fill gaps: %t\n", extractVideoArgs.FillGaps)

	// Extract video tracks
	if err := extractVideoTracks(globalArgs, extractVideoArgs, metadata, logger); err != nil {
		logger.Error("Failed to extract video tracks: %v", err)
		os.Exit(1)
	}

	logger.Info("Extract video command completed successfully")
}

func printExtractVideoUsage() {
	fmt.Fprintf(os.Stderr, "Usage: raw-tools [global options] extract-video [command options]\n\n")
	fmt.Fprintf(os.Stderr, "Generate playable video files from raw recording tracks.\n")
	fmt.Fprintf(os.Stderr, "Supports formats: webm, mp4, and others.\n\n")
	fmt.Fprintf(os.Stderr, "Command Options (Hierarchical Filtering):\n")
	fmt.Fprintf(os.Stderr, "  --userId <id>          Specify a userId or * for all (default: *)\n")
	fmt.Fprintf(os.Stderr, "  --sessionId <id>       Specify a sessionId or * for all (default: *)\n")
	fmt.Fprintf(os.Stderr, "                         Ignored if --userId=*\n")
	fmt.Fprintf(os.Stderr, "  --trackId <id>         Specify a trackId or * for all (default: *)\n")
	fmt.Fprintf(os.Stderr, "                         Ignored if --userId=* or --sessionId=*\n")
	fmt.Fprintf(os.Stderr, "  --fill_gaps            Fill with black frame when track was muted\n\n")
	fmt.Fprintf(os.Stderr, "Hierarchical Filtering Logic:\n")
	fmt.Fprintf(os.Stderr, "  --userId=*             → Extract ALL users, sessions, tracks (sessionId/trackId ignored)\n")
	fmt.Fprintf(os.Stderr, "  --userId=user1 --sessionId=*  → Extract ALL sessions/tracks for user1 (trackId ignored)\n")
	fmt.Fprintf(os.Stderr, "  --userId=user1 --sessionId=session1 --trackId=*  → Extract ALL tracks for user1/session1\n")
	fmt.Fprintf(os.Stderr, "  --userId=user1 --sessionId=session1 --trackId=track1  → Extract specific track\n\n")
	fmt.Fprintf(os.Stderr, "Examples:\n")
	fmt.Fprintf(os.Stderr, "  # Extract video for all users (sessionId/trackId ignored)\n")
	fmt.Fprintf(os.Stderr, "  raw-tools --inputFile recording.zip --output ./out extract-video --userId '*'\n\n")
	fmt.Fprintf(os.Stderr, "  # Extract video for specific user, all sessions (trackId ignored)\n")
	fmt.Fprintf(os.Stderr, "  raw-tools --inputFile recording.zip --output ./out extract-video --userId user123 --sessionId '*'\n\n")
	fmt.Fprintf(os.Stderr, "  # Extract video for specific user/session, all tracks\n")
	fmt.Fprintf(os.Stderr, "  raw-tools --inputFile recording.zip --output ./out extract-video --userId user123 --sessionId session456 --trackId '*'\n\n")
	fmt.Fprintf(os.Stderr, "Global Options: Use 'raw-tools --help' to see global options.\n")
}

func extractVideoTracks(globalArgs *GlobalArgs, extractVideoArgs *ExtractVideoArgs, metadata *RecordingMetadata, logger *getstream.DefaultLogger) error {
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

	return extractTracks(workingDir, globalArgs.Output, extractVideoArgs.UserID, extractVideoArgs.SessionID, extractVideoArgs.TrackID, metadata, "video", "both", extractVideoArgs.FillGaps, logger)
}
