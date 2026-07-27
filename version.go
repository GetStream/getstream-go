package getstream

const (
	versionName = "v5.1.1"
)

// Version returns the version of the library. versionName is written by the release
// workflow and already carries the leading "v".
func Version() string {
	return versionName
}

func versionHeader() string {
	return "stream-go-client-" + versionName
}
