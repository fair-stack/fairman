// Package download
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-01 14:41
package download

import (
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/common"
	"fmt"
	"os"
	"runtime"
)

func DownloadDocker() error {
	fmt.Println("############################################################")

	if common.PathExists("cmd/docker") {
		fmt.Println("dockerAlready downloaded，Already downloaded")
		return nil
	}

	fmt.Println("Start downloadingdocker......")

	cmd, err := common.Script("script/download_docker_binary.sh", GetOS(runtime.GOOS), GetArch(runtime.GOARCH), DockerVersion)

	if err != nil {
		return err
	}

	cmd.Path = utils.GetCurrentAbPath()

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
