package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GetStream/getstream-go/v3"
)

type MuxAVArgs struct {
	UserID    string
	SessionID string
	TrackID   string
	Media     string // "user", "display", or "both" (default)
}

type MuxAudioVideoProcess struct {
	logger *getstream.DefaultLogger
}

func NewMuxAudioVideoProcess(logger *getstream.DefaultLogger) *MuxAudioVideoProcess {
	return &MuxAudioVideoProcess{logger: logger}
}

func (p *MuxAudioVideoProcess) runMuxAV(args []string, globalArgs *GlobalArgs) {
	printHelpIfAsked(args, p.printUsage)

	// Parse command-specific flags
	fs := flag.NewFlagSet("mux-av", flag.ExitOnError)
	muxAVArgs := &MuxAVArgs{}
	fs.StringVar(&muxAVArgs.UserID, "userId", "", "Specify a userId (empty for all)")
	fs.StringVar(&muxAVArgs.SessionID, "sessionId", "", "Specify a sessionId (empty for all)")
	fs.StringVar(&muxAVArgs.TrackID, "trackId", "", "Specify a trackId (empty for all)")
	fs.StringVar(&muxAVArgs.Media, "media", "both", "Filter by media type: 'user', 'display', or 'both'")

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	// Validate input arguments against actual recording data
	metadata, err := validateInputArgs(globalArgs, muxAVArgs.UserID, muxAVArgs.SessionID, muxAVArgs.TrackID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Validation error: %v\n", err)
		os.Exit(1)
	}

	p.logger.Info("Starting mux-av command")

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
	if err := p.muxAudioVideoTracks(globalArgs, muxAVArgs, metadata, p.logger); err != nil {
		p.logger.Error("Failed to mux audio/video tracks: %v", err)
		os.Exit(1)
	}

	p.logger.Info("Mux audio and video command completed successfully")
}

func (p *MuxAudioVideoProcess) printUsage() {
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

func (p *MuxAudioVideoProcess) muxAudioVideoTracks(globalArgs *GlobalArgs, muxAVArgs *MuxAVArgs, metadata *RecordingMetadata, logger *getstream.DefaultLogger) error {
	muxer := NewAudioVideoMuxer(p.logger)
	if e := muxer.muxAudioVideoTracks(&AudioVideoMuxerConfig{
		WorkDir:     globalArgs.WorkDir,
		OutputDir:   globalArgs.Output,
		UserID:      muxAVArgs.UserID,
		SessionID:   muxAVArgs.SessionID,
		TrackID:     muxAVArgs.TrackID,
		Media:       muxAVArgs.Media,
		WithExtract: true,
		WithCleanup: false,
	}, metadata, logger); e != nil {
		return e
	}
	return nil
}
