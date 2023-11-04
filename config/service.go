// Package config
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-11 15:51
package config

type Service struct {
	Url string `mapstructure:"url" json:"url,omitempty" yaml:"url"`
}
