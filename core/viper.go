// Package core
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 14:16
package core

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/utils"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitViper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // priority: priority > priority > priority
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				if gin.Mode() == "debug" {
					config = "config_dev.yaml"
				} else if gin.Mode() == "release" {
					config = "config_pro.yaml"
				}
				fmt.Printf("You are usingconfigYou are using,configYou are using%v\n", config)
			} else {
				config = configEnv
				fmt.Printf("You are usingAPG_CONFIGYou are using,configYou are using%v\n", config)
			}
		} else {
			fmt.Printf("You are using the command line-cYou are using the command line,configYou are using the command line%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("You are usingfunc Viper()You are using,configYou are using%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.FairConf); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.FairConf); err != nil {
		fmt.Println(err)
	}
	// root Adaptability
	// according torootaccording to,according torootaccording to
	fmt.Println(filepath.Abs("."))
	global.FairConf.System.Root, _ = filepath.Abs(".")
	//global.AdpConfig.AutoCode.Root, _ = filepath.Abs("..")
	//global.BlackCache = local_cache.NewCache(
	//	local_cache.SetDefaultExpire(time.Second * time.Duration(global.AdpConfig.JWT.ExpiresTime)),
	//)
	return v
}
