package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/GetStream/getstream-go/v3"
)

type GlobalArgs struct {
	InputFile string
	InputS3   string
	Output    string
	Verbose   bool
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Parse global flags first
	globalArgs := &GlobalArgs{}
	command, remainingArgs := parseGlobalFlags(os.Args[1:], globalArgs)

	if command == "" {
		printUsage()
		os.Exit(1)
	}

	switch command {
	case "list-tracks":
		runListTracks(remainingArgs, globalArgs)
	case "extract-audio":
		runExtractAudio(remainingArgs, globalArgs)
	case "extract-video":
		runExtractVideo(remainingArgs, globalArgs)
	case "mux-av":
		runMuxAV(remainingArgs, globalArgs)
	case "mix-audio":
		runMixAudio(remainingArgs, globalArgs)
	case "process-all":
		runProcessAll(remainingArgs, globalArgs)
	case "completion":
		runCompletion(remainingArgs)
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

// parseGlobalFlags parses global flags and returns the command and remaining args
func parseGlobalFlags(args []string, globalArgs *GlobalArgs) (string, []string) {
	fs := flag.NewFlagSet("global", flag.ContinueOnError)

	fs.StringVar(&globalArgs.InputFile, "inputFile", "", "Specify raw recording zip file on file system")
	fs.StringVar(&globalArgs.InputS3, "inputS3", "", "Specify raw recording zip file on S3")
	fs.StringVar(&globalArgs.Output, "output", "", "Specify an output directory")
	fs.BoolVar(&globalArgs.Verbose, "verbose", false, "Enable verbose logging")

	// Find the command by looking for known commands
	knownCommands := map[string]bool{
		"list-tracks":   true,
		"extract-audio": true,
		"extract-video": true,
		"mux-av":        true,
		"completion":    true,
		"help":          true,
	}

	commandIndex := -1
	for i, arg := range args {
		if knownCommands[arg] {
			commandIndex = i
			break
		}
	}

	if commandIndex == -1 {
		return "", nil
	}

	// Parse global flags (everything before the command)
	globalFlags := args[:commandIndex]
	command := args[commandIndex]
	remainingArgs := args[commandIndex+1:]

	err := fs.Parse(globalFlags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing global flags: %v\n", err)
		os.Exit(1)
	}

	return command, remainingArgs
}

func setupLogger(verbose bool) *getstream.DefaultLogger {
	var level getstream.LogLevel
	if verbose {
		level = getstream.LogLevelDebug
	} else {
		level = getstream.LogLevelInfo
	}
	logger := getstream.NewDefaultLogger(os.Stderr, "", log.LstdFlags, level)
	return logger
}

func validateGlobalArgs(globalArgs *GlobalArgs, command string) error {
	if globalArgs.InputFile == "" && globalArgs.InputS3 == "" {
		return fmt.Errorf("either --inputFile or --inputS3 must be specified")
	}

	if globalArgs.InputFile != "" && globalArgs.InputS3 != "" {
		return fmt.Errorf("cannot specify both --inputFile and --inputS3")
	}

	// --output is optional for list-tracks command (it only displays information)
	if command != "list-tracks" && globalArgs.Output == "" {
		return fmt.Errorf("--output directory must be specified")
	}

	return nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Raw Recording Post Processing Tools\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [global options] <command> [command options]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Global Options:\n")
	fmt.Fprintf(os.Stderr, "  --inputFile <path>     Specify raw recording zip file on file system\n")
	fmt.Fprintf(os.Stderr, "  --inputS3 <path>       Specify raw recording zip file on S3\n")
	fmt.Fprintf(os.Stderr, "  --output <dir>         Specify an output directory (optional for list-tracks)\n")
	fmt.Fprintf(os.Stderr, "  --verbose              Enable verbose logging\n\n")
	fmt.Fprintf(os.Stderr, "Commands:\n")
	fmt.Fprintf(os.Stderr, "  list-tracks            Return list of userId - sessionId - trackId - trackType\n")
	fmt.Fprintf(os.Stderr, "  extract-audio          Generate a playable audio file (webm, mp3, ...)\n")
	fmt.Fprintf(os.Stderr, "  extract-video          Generate a playable video file (webm, mp4, ...)\n")
	fmt.Fprintf(os.Stderr, "  mux-av                 Mux audio and video tracks\n")
	fmt.Fprintf(os.Stderr, "  mix-audio              Mix multiple audio tracks into one file\n")
	fmt.Fprintf(os.Stderr, "  process-all            Process audio, video, and mux (all-in-one)\n")
	fmt.Fprintf(os.Stderr, "  completion             Generate shell completion scripts\n")
	fmt.Fprintf(os.Stderr, "  help                   Show this help message\n\n")
	fmt.Fprintf(os.Stderr, "Examples:\n")
	fmt.Fprintf(os.Stderr, "  %s --inputFile recording.zip list-tracks\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s --inputFile recording.zip --output ./out extract-audio --userId user123\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s --inputFile recording.zip --output ./out mix-audio --userId '*'\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s --verbose --inputFile recording.zip --output ./out mux-av --userId '*'\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Use '%s [global options] <command> --help' for command-specific options.\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nCompletion Setup:\n")
	fmt.Fprintf(os.Stderr, "  # Bash\n")
	fmt.Fprintf(os.Stderr, "  source <(%s completion bash)\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  # Zsh\n")
	fmt.Fprintf(os.Stderr, "  source <(%s completion zsh)\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  # Fish\n")
	fmt.Fprintf(os.Stderr, "  %s completion fish | source\n", os.Args[0])
}

func runCompletion(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: raw-tools completion <shell>\n")
		fmt.Fprintf(os.Stderr, "Supported shells: bash, zsh, fish\n")
		os.Exit(1)
	}

	shell := args[0]
	generateCompletion(shell)
}

// getInputPath returns the input path from global args
func getInputPath(globalArgs *GlobalArgs) string {
	if globalArgs.InputFile != "" {
		return globalArgs.InputFile
	}
	if globalArgs.InputS3 != "" {
		return globalArgs.InputS3
	}
	return ""
}
