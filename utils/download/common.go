// Package download
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-10 15:37
package download

var (
	DockerVersion        = "20.10.10"
	DockerComposeVersion = "1.28.2"
	HelmVersion          = "v3.6.3"
	KomposeVersion       = "v1.26.1"
	KubectlVersion       = "v1.24.0"
)

// GetOS Obtain operating system
func GetOS(goos string) string {
	switch goos {
	case "windows":
		return "win"
	case "darwin":
		return "mac"
	default:
		return goos
	}
}

// GetArch Obtain System Version
func GetArch(goarch string) string {
	switch goarch {
	case "amd64":
		return "x86_64"
	case "arm":
		return "armhf"
	case "arm64":
		return "aarch64"
	default:
		return goarch
	}
}
