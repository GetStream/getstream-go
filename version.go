package getstream

const (
	versionName = "0.0.1"
)

// Version returns the version of the library.
func Version() string {
	return "v" + versionName
}

func versionHeader() string {
	return "stream-go-client-" + versionName
}
