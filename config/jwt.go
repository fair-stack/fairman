// Package config
// @program: fairman-server-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-08 17:35
package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwtautograph
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // Expiration time
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // Buffer time
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                  // Issued by
}
