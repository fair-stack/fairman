// Package system
// @program: market-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-07-25 19:13
package system

import "time"

type SysAppVersion struct {
	Id       string    `bson:"_id"`
	CreateAt time.Time `bson:"createAt"`
	UpdateAt time.Time `bson:"updateAt"`

	// Application version
	Version string `json:"version,omitempty" form:"version,omitempty" query:"version,omitempty" bson:"version"`
	// Application Version Description
	VersionDesc string `json:"versionDesc,omitempty" form:"versionDesc,omitempty" query:"versionDesc,omitempty" bson:"versionDesc"`
	// Application version status(Application version status)
	VersionStatus string `json:"versionStatus,omitempty" form:"versionStatus,omitempty" query:"versionStatus,omitempty" bson:"versionStatus"`
	// Backend software package
	Backend SysBaseFiles `json:"backend,omitempty" form:"backend,omitempty" query:"backend,omitempty" bson:"backend"`
	// Front end software package
	Frontend SysBaseFiles `json:"frontend,omitempty" form:"frontend,omitempty" query:"frontend,omitempty" bson:"frontend"`
}
