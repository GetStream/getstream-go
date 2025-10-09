package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/GetStream/getstream-go/v3"
	"github.com/GetStream/getstream-go/v3/cmd/raw-recording-tools/webm"
)

// Generic track extraction function that works for both audio and video
func extractTracks(globalArgs *GlobalArgs, userID, sessionID, trackID, trackType, mediaFilter string, fillGaps bool, logger *getstream.DefaultLogger) error {
	var inputPath string
	if globalArgs.InputFile != "" {
		inputPath = globalArgs.InputFile
	} else {
		// TODO: Handle S3 input
		return fmt.Errorf("S3 input not yet supported")
	}

	// Extract to temp directory if needed (unified approach)
	workingDir, cleanup, err := extractToTempDir(inputPath, logger)
	if err != nil {
		return fmt.Errorf("failed to prepare working directory: %w", err)
	}
	defer cleanup()

	// Parse metadata from the working directory
	parser := NewMetadataParser(logger)
	metadata, err := parser.ParseRecording(workingDir)
	if err != nil {
		return fmt.Errorf("failed to parse recording metadata: %w", err)
	}

	// Filter tracks to specified type only and apply hierarchical filtering
	filteredTracks := parser.FilterTracks(metadata.Tracks, userID, sessionID, trackID)
	typedTracks := make([]*TrackInfo, 0)
	for _, track := range filteredTracks {
		if track.TrackType == trackType {
			// Apply media type filtering if specified
			if mediaFilter != "" && mediaFilter != "both" {
				if mediaFilter == "user" && track.IsScreenshare {
					continue // Skip display tracks when only user requested
				}
				if mediaFilter == "display" && !track.IsScreenshare {
					continue // Skip user tracks when only display requested
				}
			}
			typedTracks = append(typedTracks, track)
		}
	}

	if len(typedTracks) == 0 {
		logger.Warn("No %s tracks found matching the filter criteria", trackType)
		return nil
	}

	logger.Info("Found %d %s tracks to extract", len(typedTracks), trackType)

	// Create output directory if it doesn't exist
	err = os.MkdirAll(globalArgs.Output, 0755)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Extract and convert each track
	for i, track := range typedTracks {
		logger.Info("Processing %s track %d/%d: %s", trackType, i+1, len(typedTracks), track.TrackID)

		err = extractSingleTrackWithOptions(workingDir, track, globalArgs.Output, trackType, fillGaps, logger)
		if err != nil {
			logger.Error("Failed to extract %s track %s: %v", trackType, track.TrackID, err)
			continue
		}
	}

	return nil
}

func extractSingleTrackWithOptions(inputPath string, track *TrackInfo, outputDir string, trackType string, fillGaps bool, logger *getstream.DefaultLogger) error {
	// Create a temp directory for extraction and processing
	tempDir, err := os.MkdirTemp("", fmt.Sprintf("%s-extract-*", trackType))
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy track files from working directory (now always a directory)
	err = copyTrackFiles(inputPath, track, tempDir, trackType)
	if err != nil {
		return fmt.Errorf("failed to copy track files: %w", err)
	}

	// Convert using the WebM converter
	err = webm.ConvertDirectory(tempDir, logger)
	if err != nil {
		return fmt.Errorf("failed to convert %s track: %w", trackType, err)
	}

	// Find ALL generated .webm files
	webmFiles, _ := filepath.Glob(filepath.Join(tempDir, "*.webm"))
	if len(webmFiles) == 0 {
		return fmt.Errorf("no webm output files found")
	}

	logger.Info("Found %d WebM segment files for %s track %s", len(webmFiles), trackType, track.TrackID)

	// Create segments with timing info and fill gaps
	finalFile, err := processSegmentsWithGapFilling(webmFiles, track, trackType, outputDir, fillGaps, logger)
	if err != nil {
		return fmt.Errorf("failed to process segments with gap filling: %w", err)
	}

	logger.Info("Successfully extracted %s track to: %s", trackType, finalFile)
	return nil
}

// NOTE: extractTrackFiles removed - now always use copyTrackFiles since we always work with directories

// copyTrackFiles copies the rtpdump and sdp files for a specific track to the destination directory
func copyTrackFiles(inputPath string, track *TrackInfo, destDir string, trackType string) error {
	// Walk through the input directory and copy files related to this track
	return filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileName := info.Name()
			// Check if this file belongs to our track (using trackType parameter)
			if strings.Contains(fileName, track.TrackID) && strings.Contains(fileName, trackType) {
				if strings.HasSuffix(fileName, ".rtpdump") || strings.HasSuffix(fileName, ".sdp") {
					// Copy this file to destination
					destPath := filepath.Join(destDir, fileName)

					err = copyFile(path, destPath)
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
}

// Helper function to copy a file
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = srcFile.WriteTo(dstFile)
	return err
}

// processSegmentsWithGapFilling processes webm segments, fills gaps if requested, and concatenates into final file
func processSegmentsWithGapFilling(webmFiles []string, track *TrackInfo, trackType string, outputDir string, fillGaps bool, logger *getstream.DefaultLogger) (string, error) {
	if len(webmFiles) == 1 {
		// Single segment, just copy it with final name
		finalName := fmt.Sprintf("%s_%s_%s_%s.webm", trackType, track.UserID, track.SessionID, track.TrackID)
		finalPath := filepath.Join(outputDir, finalName)
		err := copyFile(webmFiles[0], finalPath)
		return finalPath, err
	}

	// Multiple segments - sort, optionally fill gaps, and concatenate
	if fillGaps {
		logger.Info("Processing %d segments with gap filling for %s track %s", len(webmFiles), trackType, track.TrackID)
	} else {
		logger.Info("Processing %d segments (no gap filling) for %s track %s", len(webmFiles), trackType, track.TrackID)
	}

	// Map webm files to their original segment timing using filenames
	segmentMap := make(map[string]string) // originalFilename -> webmFilePath
	for _, webmFile := range webmFiles {
		// Extract original filename from webm filename (remove .webm, add .rtpdump)
		baseName := strings.TrimSuffix(filepath.Base(webmFile), ".webm")
		originalName := baseName + ".rtpdump"
		segmentMap[originalName] = webmFile
	}

	// Build list of files to concatenate (with optional gap fillers)
	var filesToConcat []string

	for i, segment := range track.Segments {
		// Add the segment file
		filesToConcat = append(filesToConcat, segmentMap[segment.BaseFilename])

		// Add gap filler if requested and there's a gap before the next segment
		if fillGaps && i < track.SegmentCount-1 {
			nextSegment := track.Segments[i+1]
			gapDuration := FirstPacketNtpTimestamp(nextSegment) - LastPacketNtpTimestamp(segment)

			if gapDuration > 0 { // There's a gap
				gapSeconds := float64(gapDuration) / 1000.0
				logger.Info("Detected %dms gap between segments, generating %s filler", gapDuration, trackType)

				// Create gap filler file
				gapFilePath := filepath.Join(outputDir, fmt.Sprintf("gap_%s_%d.webm", trackType, i))

				if trackType == "audio" {
					err := webm.GenerateSilence(gapFilePath, gapSeconds, logger)
					if err != nil {
						logger.Warn("Failed to generate silence, skipping gap: %v", err)
						continue
					}
				} else if trackType == "video" {
					// Use VP8 codec and 720p quality as defaults
					err := webm.GenerateBlackVideo(gapFilePath, "video/VP8", gapSeconds, 1280, 720, 30, logger)
					if err != nil {
						logger.Warn("Failed to generate black video, skipping gap: %v", err)
						continue
					}
				}

				filesToConcat = append(filesToConcat, gapFilePath)
			}
		}
	}

	// Create final output file
	finalName := fmt.Sprintf("%s_%s_%s_%s.webm", trackType, track.UserID, track.SessionID, track.TrackID)
	finalPath := filepath.Join(outputDir, finalName)

	// Concatenate all segments (with gap fillers if any)
	err := webm.ConcatFile(finalPath, filesToConcat, logger)
	if err != nil {
		return "", fmt.Errorf("failed to concatenate segments: %w", err)
	}

	// Clean up temporary gap filler files
	if fillGaps {
		for _, file := range filesToConcat {
			if strings.Contains(file, "gap_") {
				os.Remove(file)
			}
		}
	}

	if fillGaps {
		logger.Info("Successfully concatenated %d segments with gap filling into %s", track.SegmentCount, finalPath)
	} else {
		logger.Info("Successfully concatenated %d segments into %s", track.SegmentCount, finalPath)
	}
	return finalPath, nil
}

// extractToTempDir extracts archive to temp directory or returns the directory path
// Returns: (workingDir, cleanupFunc, error)
func extractToTempDir(inputPath string, logger *getstream.DefaultLogger) (string, func(), error) {
	// If it's already a directory, just return it
	if stat, err := os.Stat(inputPath); err == nil && stat.IsDir() {
		logger.Debug("Input is already a directory: %s", inputPath)
		return inputPath, func() {}, nil
	}

	// If it's a tar.gz file, extract it to temp directory
	if strings.HasSuffix(strings.ToLower(inputPath), ".tar.gz") {
		logger.Info("Extracting tar.gz archive to temporary directory...")

		tempDir, err := os.MkdirTemp("", "raw-tools-*")
		if err != nil {
			return "", nil, fmt.Errorf("failed to create temp directory: %w", err)
		}

		cleanup := func() {
			os.RemoveAll(tempDir)
		}

		err = extractTarGzToDir(inputPath, tempDir, logger)
		if err != nil {
			cleanup()
			return "", nil, fmt.Errorf("failed to extract tar.gz: %w", err)
		}

		logger.Debug("Extracted archive to: %s", tempDir)
		return tempDir, cleanup, nil
	}

	return "", nil, fmt.Errorf("unsupported input format: %s (only tar.gz files and directories supported)", inputPath)
}

// extractTarGzToDir extracts a tar.gz file to the specified directory
func extractTarGzToDir(tarGzPath, destDir string, logger *getstream.DefaultLogger) error {
	file, err := os.Open(tarGzPath)
	if err != nil {
		return fmt.Errorf("failed to open tar.gz file: %w", err)
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar entry: %w", err)
		}

		// Skip directories
		if header.FileInfo().IsDir() {
			continue
		}

		// Create destination file
		destPath := filepath.Join(destDir, header.Name)

		// Create directory structure if needed
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return fmt.Errorf("failed to create directory structure: %w", err)
		}

		// Extract file
		outFile, err := os.Create(destPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", destPath, err)
		}

		_, err = io.Copy(outFile, tarReader)
		outFile.Close()
		if err != nil {
			return fmt.Errorf("failed to extract file %s: %w", destPath, err)
		}
	}

	return nil
}
