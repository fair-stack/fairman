// Package global
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 14:15
package global

import (
	"cnic/fairman-gin/config"

	"golang.org/x/sync/singleflight"

	ut "github.com/go-playground/universal-translator"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

var (
	FairVp        *viper.Viper
	FairConf      config.Server
	FairLog       *zap.Logger
	FairValidator ut.Translator

	FairClient             ClientFactory
	FairFileSystem         FileSystemService
	FairConcurrencyControl = &singleflight.Group{}

	//FairDB    typings.Connection
	//FairStore *datastore.Store
)
