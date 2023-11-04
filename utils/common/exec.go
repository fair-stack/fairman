// Package common
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-10 14:14
package common

import (
	"fmt"
	"os"
	"os/exec"
)

// Script returns a command to execute a script through a shell.
func Script(args ...string) (*exec.Cmd, error) {
	shell := "/bin/sh"
	if other := os.Getenv("SHELL"); other != "" {
		if other != "/bin/zsh" {
			shell = other
		}
	}
	fmt.Println(shell, args)
	return exec.Command(shell, args...), nil
}

// PathExists Determine if a file or folder exists
// Enter file path，Enter file pathboolEnter file path
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func MV(path1, path2 string) error {
	var cmd *exec.Cmd
	cmd = exec.Command("mv", path1, path2)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return err
}
