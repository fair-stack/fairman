// Package download
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-10 15:43
package download

import (
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/common"
	"fmt"
	"os"
	"runtime"
)

func DownloadKompose() error {
	fmt.Println("############################################################")

	if common.PathExists("cmd/kompose") {
		fmt.Println("komposeAlready downloaded，Already downloaded")
		return nil
	}

	cmd, err := common.Script("script/download_kompose_binary.sh", runtime.GOOS, runtime.GOARCH, KomposeVersion)

	if err != nil {
		return err
	}

	cmd.Dir = utils.GetCurrentAbPath()
	fmt.Println("Start downloadingkompose......")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
