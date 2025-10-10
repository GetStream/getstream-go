package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/GetStream/getstream-go/v3"
	"github.com/GetStream/getstream-go/v3/cmd/raw-recording-tools/webm"
)

// MixAudioArgs represents the arguments for the mix-audio command
type MixAudioArgs struct {
	IncludeScreenShare bool
}

// AudioFileWithTiming represents an audio file with its timing information
type AudioFileWithTiming struct {
	FilePath      string     // Path to the WebM audio file
	UserID        string     // User who created this audio
	SessionID     string     // Session ID
	TrackID       string     // Track ID
	StartOffsetMs int64      // When this audio should start in the final mix (milliseconds)
	DurationMs    int64      // Duration of the audio file
	EndOffsetMs   int64      // When this audio ends (StartOffsetMs + DurationMs)
	TrackInfo     *TrackInfo // Original track metadata
}

type MixAudioProcess struct {
	logger *getstream.DefaultLogger
}

func NewMixAudioProcess(logger *getstream.DefaultLogger) *MixAudioProcess {
	return &MixAudioProcess{logger: logger}
}

// runMixAudio handles the mix-audio command
func (p *MixAudioProcess) runMixAudio(args []string, globalArgs *GlobalArgs) {
	printHelpIfAsked(args, p.printUsage)

	mixAudioArgs := &MixAudioArgs{
		IncludeScreenShare: false,
	}

	// Validate input arguments against actual recording data
	metadata, err := validateInputArgs(globalArgs, "", "", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Validation error: %v\n", err)
		os.Exit(1)
	}

	p.logger.Info("Starting mix-audio command")

	// Execute the mix-audio operation
	if e := p.mixAllAudioTracks(globalArgs, mixAudioArgs, metadata, p.logger); e != nil {
		p.logger.Error("Mix-audio failed: %v", e)
		os.Exit(1)
	}

	p.logger.Info("Mix-audio command completed successfully")
}

// mixAllAudioTracks orchestrates the entire audio mixing workflow using existing extraction logic
func (p *MixAudioProcess) mixAllAudioTracks(globalArgs *GlobalArgs, mixAudioArgs *MixAudioArgs, metadata *RecordingMetadata, logger *getstream.DefaultLogger) error {
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

	// Step 2: Find all extracted audio files and prepare them for mixing
	logger.Info("Step 2/2: Discovering extracted files and mixing...")
	audioFiles, err := p.discoverExtractedAudioFiles(globalArgs.Output, logger)
	if err != nil {
		return fmt.Errorf("failed to discover extracted audio files: %w", err)
	}

	if len(audioFiles) == 0 {
		return fmt.Errorf("no audio files were extracted - check your filter criteria")
	}

	logger.Info("Found %d extracted audio files to mix", len(audioFiles))

	// Step 3: Mix all discovered audio files using existing webm.MixAudioFiles
	outputFile := filepath.Join(globalArgs.Output, "mixed_audio.webm")

	// Convert AudioFileWithTiming to the format expected by webm.MixAudioFiles
	// webm.MixAudioFiles expects: map[string]int where key=filepath, value=offset_ms
	fileOffsetMap := make(map[string]int)
	for _, audioFile := range audioFiles {
		fileOffsetMap[audioFile.FilePath] = int(audioFile.StartOffsetMs)
	}

	err = webm.MixAudioFiles(outputFile, fileOffsetMap, logger)
	if err != nil {
		return fmt.Errorf("failed to mix audio files: %w", err)
	}

	logger.Info("Successfully created mixed audio file: %s", outputFile)

	// Clean up individual audio files (optional)
	for _, audioFile := range audioFiles {
		if err := os.Remove(audioFile.FilePath); err != nil {
			logger.Warn("Failed to clean up temporary file %s: %v", audioFile.FilePath, err)
		}
	}

	return nil
}

// discoverExtractedAudioFiles finds all audio files that were extracted and prepares them for mixing
func (p *MixAudioProcess) discoverExtractedAudioFiles(outputDir string, logger *getstream.DefaultLogger) ([]AudioFileWithTiming, error) {
	var audioFiles []AudioFileWithTiming

	// Find all .webm audio files in the output directory
	err := filepath.Walk(outputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Look for audio WebM files (created by extractTracks)
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".webm") && strings.Contains(strings.ToLower(info.Name()), "audio") {
			logger.Debug("Found extracted audio file: %s", path)

			// Parse filename to extract timing info (if available)
			audioFile := AudioFileWithTiming{
				FilePath:      path,
				StartOffsetMs: 0, // Will be calculated from metadata if needed
				DurationMs:    0, // Will be calculated if needed
				EndOffsetMs:   0, // Will be calculated if needed
			}

			// Try to extract user/session/track info from filename
			// Expected format: audio_userID_sessionID_trackID.webm
			basename := filepath.Base(path)
			basename = strings.TrimSuffix(basename, ".webm")
			parts := strings.Split(basename, "_")

			if len(parts) >= 4 && parts[0] == "audio" {
				audioFile.UserID = parts[1]
				audioFile.SessionID = parts[2]
				audioFile.TrackID = parts[3]
			}

			audioFiles = append(audioFiles, audioFile)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan output directory: %w", err)
	}

	// Sort by filename for consistent ordering
	sort.Slice(audioFiles, func(i, j int) bool {
		return audioFiles[i].FilePath < audioFiles[j].FilePath
	})

	return audioFiles, nil
}

// Note: We removed mixAudioFilesUsingExistingLogic since we now use webm.MixAudioFiles directly

// printMixAudioUsage prints the usage information for the mix-audio command
func (p *MixAudioProcess) printUsage() {
	fmt.Println("Usage: raw-tools [global-options] mix-audio [options]")
	fmt.Println()
	fmt.Println("Mix all audio tracks from multiple users/sessions into a single audio file")
	fmt.Println("with proper timing synchronization (like a conference call recording).")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --userId <id>        Filter by user ID (* for all users, default: *)")
	fmt.Println("  --sessionId <id>     Filter by session ID (* for all sessions, default: *)")
	fmt.Println("  --trackId <id>       Filter by track ID (* for all tracks, default: *)")
	fmt.Println("  --no-fill-gaps       Don't fill gaps with silence (not recommended for mixing)")
	fmt.Println("  -h, --help           Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  # Mix all audio tracks from all users and sessions")
	fmt.Println("  raw-tools --inputFile recording.tar.gz --output /tmp/mixed mix-audio")
	fmt.Println()
	fmt.Println("  # Mix audio tracks from a specific user")
	fmt.Println("  raw-tools --inputFile recording.tar.gz --output /tmp/mixed mix-audio --userId user123")
	fmt.Println()
	fmt.Println("  # Mix audio tracks from a specific session")
	fmt.Println("  raw-tools --inputFile recording.tar.gz --output /tmp/mixed mix-audio --sessionId session456")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  Creates 'mixed_audio.webm' - a single audio file containing all mixed tracks")
	fmt.Println("  with proper timing synchronization based on the original recording timeline.")
}
