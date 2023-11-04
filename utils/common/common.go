// Package common
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2023-07-05 18:09
package common

import (
	"fmt"
	"net"
	"strconv"
)

func ScanPort(port int, network string) bool {
	fmt.Println(port)
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer ln.Close()
	return false
}
