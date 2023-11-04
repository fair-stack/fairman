// Package utils
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-13 17:24
package utils

import (
	"cnic/fairman-gin/global"
	"fmt"
	"runtime"
)

func Errorf(any interface{}, opt ...interface{}) error {
	if any != nil {
		err := (error)(nil)

		switch any.(type) {
		case string:
			err = fmt.Errorf(any.(string), opt...)
		case error:
			err = fmt.Errorf(any.(error).Error(), opt...)
		default:
			err = fmt.Errorf("%v", err)
		}

		_, fn, line, _ := runtime.Caller(1)
		global.FairLog.Error(fmt.Sprintf("[%s:%d] %v \n", fn, line, err))
		return err
	}
	return nil
}
