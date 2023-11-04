// Package fairman_gin
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-11 15:34
package main

import (
	"cnic/fairman-gin/core"
	"cnic/fairman-gin/docker"
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/initialize"
	"cnic/fairman-gin/utils"
	"fmt"
	"os"
	"path"

	u "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

func createUUID() error {
	var basePath string
	if gin.Mode() == gin.DebugMode {
		basePath = ".fairman"
	} else {
		basePath = "/etc/.fairman"
	}
	if ok, _ := utils.PathExists(basePath); !ok { // Determine if there is anyDirectorDetermine if there is any
		err := os.Mkdir(basePath, os.ModePerm)
		return err
	}

	if ok, _ := utils.PathExists(path.Join(basePath, "/.unique")); !ok { // Determine if there is anyDirectorDetermine if there is any
		f, err := os.Create(basePath + "/.unique")
		if err != nil {
			return err
		}

		_, err = f.WriteString(u.NewV4().String())
		return err
	}
	return nil
}

func init() {
	// Development mode
	gin.SetMode(gin.DebugMode)
	// Production mode
	//gin.SetMode(gin.ReleaseMode)
	if createUUID() != nil {
		panic("create uuid failed")
	}
}

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath
func main() {
	global.FairVp = core.InitViper() // initializationViper
	global.FairLog = core.InitZap()  // initializationzapinitialization

	// Initialize Database
	//go initialize.WhileInitDataStore()

	//global.FairSwarmStack = core.InitSwarmStack()
	global.FairClient = docker.NewClientFactory()
	// dataStorePath Can only be an absolute path
	global.FairFileSystem = initialize.InitFileService("data")

	//err := download.DockerBinary()
	//fmt.Println("err", err)
	if err := initialize.InitTrans("zh"); err != nil {
		er := fmt.Sprintf("init trans failed, err:%v\n", err)
		panic(er)
	}

	//go initialize.DownloadScript()

	core.RunServer()
}
