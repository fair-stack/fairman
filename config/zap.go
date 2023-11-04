// Package config
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 14:42
package config

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // level
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // output
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // Log Prefix
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // Log folder
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`                 // Show Rows
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // Encoding level
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // Stack name
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // Output Console
}
