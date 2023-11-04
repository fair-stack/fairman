// Package initialize
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-24 17:28
package initialize

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/utils/filesystem"
	"log"
)

func InitFileService(dataStorePath string) global.FileSystemService {
	fileService, err := filesystem.NewService(dataStorePath, "")
	if err != nil {
		log.Fatalf("failed creating file service: %v", err)
	}
	return fileService
}
