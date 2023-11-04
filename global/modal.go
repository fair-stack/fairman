// Package global
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-09-19 15:01
package global

import (
	"github.com/pquerna/ffjson/ffjson"
)

type ChanLogModal struct {
	// news
	Message string `json:"message,omitempty"`
	// state
	Level string `json:"level,omitempty"`
	// schedule
	Progress float64 `json:"progress,omitempty"`
	// Steps
	Steps int `json:"steps,omitempty"`
}

func (c *ChanLogModal) ToJsonByte() ([]byte, error) {
	res, err := ffjson.Marshal(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ChanLogModal) ToJsonStr() (string, error) {
	b, err := c.ToJsonByte()

	return string(b), err
}
