package getstream

const (
	versionName = "v3.1.0-feeds"
)

// Version returns the version of the library.
func Version() string {
	return "v" + versionName
}

func versionHeader() string {
	return "stream-go-client-" + versionName
}
