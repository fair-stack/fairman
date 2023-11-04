// Package initialize
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-10 16:00
package initialize

import (
	"cnic/fairman-gin/utils/download"
	"fmt"
	"os"
	"time"
)

func DownloadScript() {
	err := os.MkdirAll("cmd", 0755)
	if err != nil {
		return
	}

	time.Sleep(1 * time.Second)
	fmt.Println()

	err = download.DownloadDocker()
	if err != nil {
		os.Remove("cmd/docker")
		panic(err)
	}

	err = download.DownloadDockerCompose()
	if err != nil {
		os.Remove("cmd/docker-compose")
		panic(err)
	}

	// implementswarm init
	//exec.Command("cmd/docker", "swarm", "init").Run()

	err = download.DownloadKubectl()
	if err != nil {
		os.Remove("cmd/kubectl")
		panic(err)
	}

	err = download.DownloadKompose()
	if err != nil {
		os.Remove("cmd/kompose")
		panic(err)
	}

	err = download.DownloadHelm()
	if err != nil {
		os.Remove("cmd/helm")
		panic(err)
	}

	fmt.Println("All downloads complete")
}
