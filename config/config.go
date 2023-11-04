// Package config
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 14:20
package config

type Server struct {
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	RateLimit RateLimit `mapstructure:"rate-limit" json:"rate-limit" yaml:"rate-limit"`
	Docker    Docker    `mapstructure:"docker" json:"docker" yaml:"docker"`
	Service   Service   `mapstructure:"service" json:"service,omitempty" yaml:"service"`
}

//var (
//	once       sync.Once
//	configFile = flag.String("config", "", "config file")
//)
//
//func init() {
//	flag.Parse()
//	once.Do(initConfig)
//}
//
//func initConfig() {
//	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
//	fmt.Println(*configFile)
//	global.FairVp = core.InitViper(*configFile)
//}
