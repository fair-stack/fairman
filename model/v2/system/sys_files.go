// Package system
// @program: market-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-07-20 17:40
package system

import "time"

type SysBaseFiles struct {
	// Default parameters _id Default parameters
	Id       string    `bson:"_id"`
	CreateAt time.Time `bson:"createAt"`
	UpdateAt time.Time `bson:"updateAt"`
	// url address
	Url string `json:"url,omitempty" form:"url,omitempty" query:"url,omitempty" bson:"url"`
	// file name
	Name string `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name"`
	// original file name
	OriginalName string `json:"originalName,omitempty" form:"originalName,omitempty" query:"originalName,omitempty" bson:"originalName"`
	// file size
	Size int64 `json:"size,omitempty" form:"size,omitempty" query:"size,omitempty" bson:"size"`
	// file extension
	Suffix string `json:"suffix,omitempty" form:"suffix,omitempty" query:"suffix,omitempty" bson:"suffix"`
}
