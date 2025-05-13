package version

var version string

func Get() string {
	version = "dev"
	return version
}
