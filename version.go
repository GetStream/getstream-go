package getstream

const (
	versionName = "1.1.0"
)

// Version returns the version of the library.
func Version() string {
	return "v" + versionName
}

func versionHeader() string {
	return "stream-go-client-" + versionName
}
