// Package utils
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-15 15:39
package utils

import (
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/pkg/errors"
)

func osProgram(program string) string {
	if runtime.GOOS == "windows" {
		program += ".exe"
	}
	return program
}

func ProgramPath(rootPath, program string) string {
	return path.Join(rootPath, osProgram(program))
}

// IsBinaryPresent check if docker compose binary is present
func IsBinaryPresent(program string) bool {
	_, err := exec.LookPath(program)
	return err == nil
}

// Copy copies sourcePath to destinationPath
func Copy(sourcePath, destinationPath string) error {
	si, err := os.Stat(sourcePath)
	if err != nil {
		return errors.WithMessage(err, "file check failed")
	}

	input, err := os.ReadFile(sourcePath)
	if err != nil {
		return errors.WithMessage(err, "failed reading file")
	}

	err = os.WriteFile(destinationPath, input, si.Mode())
	if err != nil {
		return errors.WithMessage(err, "failed writing file")
	}

	return nil
}
