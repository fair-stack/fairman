// Package errno
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 17:48
package errno

import "github.com/pkg/errors"

var (
	// OK = &Errno{Code: 0, Message: "OK"}

	// TokenExpired token Expired
	TokenExpired = errors.New("token is expired")
)

const (
	ERROR   = -1
	SUCCESS = 0
	TIMEOUT = 10001
)
