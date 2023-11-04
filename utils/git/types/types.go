// Package types
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-17 17:39
package types

// RepoConfig represents a configuration for a repo
type RepoConfig struct {
	// The repo url
	URL string `example:""`
	// The reference name
	ReferenceName string `example:""`
	// Path to where the config file is in this url/refName
	ConfigFilePath string `example:""`
	// Git credentials
	Authentication *GitAuthentication
	// Repository hash
	ConfigHash string `example:""`
}

type GitAuthentication struct {
	Username string
	Password string
}
