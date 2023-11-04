// Package download
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-10 15:41
package download

import (
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/common"
	"fmt"
	"os"
	"runtime"
)

func DownloadDockerCompose() error {
	if common.PathExists("cmd/docker-compose") {
		fmt.Println("docker-composeAlready downloaded，Already downloaded")
		return nil
	}

	cmd, err := common.Script("script/download_docker_compose_binary.sh", GetOS(runtime.GOOS), GetArch(runtime.GOARCH), DockerComposeVersion)

	if err != nil {
		return err
	}
	cmd.Dir = utils.GetCurrentAbPath()

	fmt.Println("Start downloadingdocker-compose......")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
