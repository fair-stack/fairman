// Package config
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 17:58
package config

type RateLimit struct {
	IpVerify   bool   `mapstructure:"ip-verify" json:"ip-verify" yaml:"ip-verify"`          // Open or notipOpen or not
	IpLimitCon int64  `mapstructure:"ip-limit-con" json:"ip-limit-con" yaml:"ip-limit-con"` // Access per secondipAccess per second
	IpListKey  string `mapstructure:"ip-list-key" json:"ip-list-key" yaml:"ip-list-key"`    // ipList ofkey

	Cap     int64 `mapstructure:"cap" json:"cap" yaml:"cap"`             // Initialize quantity
	Quantum int64 `mapstructure:"quantum" json:"quantum" yaml:"quantum"` // Increase quantity per second

	LimitCountIP int `mapstructure:"iplimit-count" json:"iplimitCount" yaml:"iplimit-count"` // IPLimit the number of times LimitTimeIP Limit the number of times Limit the number of times LimitCountIP Limit the number of times
	LimitTimeIP  int `mapstructure:"iplimit-time" json:"iplimitTime" yaml:"iplimit-time"`    // IPLimit for one hour Limit for one hour 3600
}
