// Package download
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-10 15:45
package download

import (
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/common"
	"fmt"
	"os"
	"runtime"
)

func DownloadKubectl() error {
	fmt.Println("############################################################")

	if common.PathExists("cmd/kubectl") {
		fmt.Println("kubectlAlready downloaded，Already downloaded")
		return nil
	}

	cmd, err := common.Script("script/download_kubectl_binary.sh", runtime.GOOS, runtime.GOARCH, KubectlVersion)

	if err != nil {
		return err
	}

	cmd.Dir = utils.GetCurrentAbPath()
	fmt.Println("Start downloadingkubectl......")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
