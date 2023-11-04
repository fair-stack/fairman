// Package utils
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 17:45
package utils

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/pkg/errors"
)

func Reload() error {
	if runtime.GOOS == "windows" {
		return errors.New("System does not support")
	}
	pid := os.Getpid()
	cmd := exec.Command("kill", "-1", strconv.Itoa(pid))
	return cmd.Run()
}
