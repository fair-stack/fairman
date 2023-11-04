// Package stackutils
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-17 17:35
package stackutils

import (
	"cnic/fairman-gin/model/v2/docker"
	"cnic/fairman-gin/utils/filesystem"
)

// GetStackFilePaths returns a list of file paths based on stack project path
// Returns a list of file paths based on stack item paths
func GetStackFilePaths(stack docker.DocSoftware) []string {
	var filePaths []string
	//for _, file := range append([]string{stack.EntryPoint}, stack.AdditionalFiles...) {
	for _, file := range append([]string{stack.EntryPoint}) {
		filePaths = append(filePaths, filesystem.JoinPaths(stack.ProjectPath, file))
	}
	return filePaths
}
