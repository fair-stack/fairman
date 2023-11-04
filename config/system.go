// Package config
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 14:50
package config

type System struct {
	Name         string `mapstructure:"name" json:"name" yaml:"name"`                           // entry name
	Version      string `mapstructure:"version" json:"version" yaml:"version"`                  // Project version
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                           // Port value
	Login        string `mapstructure:"login" json:"login" yaml:"login"`                        // Login interface address
	DbType       string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                   // Database type:mysql(Database type)|sqlite|sqlserver|postgresql
	PIDFile      string `mapstructure:"pid-file" json:"pidFile" yaml:"pid-file"`                // processidprocess
	Root         string `mapstructure:"root" json:"root" yaml:"root"`                           // Project Root Directory
	BackendPath  string `mapstructure:"backend-path" json:"backendPath" yaml:"backend-path"`    // Backend software package path
	FrontendPath string `mapstructure:"frontend-path" json:"frontendPath" yaml:"frontend-path"` // Frontend software package path
}
