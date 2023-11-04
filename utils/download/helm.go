// Package download
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-10 15:42
package download

import (
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/common"
	"fmt"
	"os"
	"runtime"
)

func DownloadHelm() error {
	fmt.Println("############################################################")

	if common.PathExists("cmd/helm") {
		fmt.Println("helmAlready downloaded，Already downloaded")
		return nil
	}

	cmd, err := common.Script("script/download_helm_binary.sh", runtime.GOOS, runtime.GOARCH, HelmVersion)

	if err != nil {
		return err
	}

	cmd.Dir = utils.GetCurrentAbPath()
	fmt.Println("Start downloadinghelm......")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
