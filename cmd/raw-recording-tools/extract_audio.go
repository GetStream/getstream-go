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
	fs.StringVar(&extractAudioArgs.UserID, "userId", "*", "Specify a userId or * for all")
	fs.StringVar(&extractAudioArgs.SessionID, "sessionId", "*", "Specify a sessionId or * for all")
	fs.StringVar(&extractAudioArgs.TrackID, "trackId", "*", "Specify a trackId or * for all")
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

	// Validate global arguments
	if err := validateGlobalArgs(globalArgs, "extract-audio"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		printExtractAudioUsage()
		os.Exit(1)
	}

	// Validate input arguments against actual recording data
	if err := validateInputArgs(globalArgs, extractAudioArgs.UserID, extractAudioArgs.SessionID, extractAudioArgs.TrackID); err != nil {
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

	if extractAudioArgs.UserID == "*" {
		fmt.Printf("  → Processing ALL users (sessionId/trackId ignored)\n")
	} else {
		fmt.Printf("  Session ID filter: %s\n", extractAudioArgs.SessionID)
		if extractAudioArgs.SessionID == "*" {
			fmt.Printf("  → Processing ALL sessions for user '%s' (trackId ignored)\n", extractAudioArgs.UserID)
		} else {
			fmt.Printf("  Track ID filter: %s\n", extractAudioArgs.TrackID)
			if extractAudioArgs.TrackID == "*" {
				fmt.Printf("  → Processing ALL tracks for user '%s', session '%s'\n", extractAudioArgs.UserID, extractAudioArgs.SessionID)
			} else {
				fmt.Printf("  → Processing specific track '%s' for user '%s', session '%s'\n", extractAudioArgs.TrackID, extractAudioArgs.UserID, extractAudioArgs.SessionID)
			}
		}
	}
	fmt.Printf("  Fill gaps: %t\n", extractAudioArgs.FillGaps)

	// Implement extract audio functionality
	err := extractAudioTracks(globalArgs, extractAudioArgs, logger)
	if err != nil {
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

func extractAudioTracks(globalArgs *GlobalArgs, extractAudioArgs *ExtractAudioArgs, logger *getstream.DefaultLogger) error {
	return extractTracks(globalArgs, extractAudioArgs.UserID, extractAudioArgs.SessionID, extractAudioArgs.TrackID, "audio", "both", extractAudioArgs.FillGaps, logger)
}
