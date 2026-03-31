package getstream

const (
	versionName = "v4.0.4"
)

// Version returns the version of the library.
func Version() string {
	return "v" + versionName
}

func versionHeader() string {
	return "stream-go-client-" + versionName
}
