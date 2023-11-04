// Package config
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-24 15:46
package config

type Docker struct {
	BinaryPath string `mapstructure:"binary-path" json:"binary-path" yaml:"binary-path"` // dockerWait for binary file address
	ConfigPath string `mapstructure:"config-path" json:"config-path" yaml:"config-path"` // Profile Address docker-compose.yamlProfile Address
	Type       string `mapstructure:"type" json:"type" yaml:"type"`                      // Docker Link Type local Link Typeurl agent/edge Link Typeurl
	Url        string `mapstructure:"url" json:"url" yaml:"url"`                         // agent/edge Address required （Address required）
}
