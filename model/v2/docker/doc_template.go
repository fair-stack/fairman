// Package docker
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-21 16:40
package docker

import (
	"cnic/fairman-gin/model/v2/system"
	"time"
)

type DocTemplate struct {
	// Default parameters _id Default parameters
	Id       string    `bson:"_id"`
	CreateAt time.Time `bson:"createAt"`
	UpdateAt time.Time `bson:"updateAt"`
	// Software Name
	Name string `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty" bson:"name" binding:"required"` //  binding:"required"
	// Software name in Chinese
	NameZh string `json:"nameZh,omitempty" form:"nameZh,omitempty" query:"nameZh,omitempty" bson:"nameZh"`
	// softwarelogo softwarehttpsoftware
	Logo string `json:"logo,omitempty" form:"logo,omitempty" query:"logo,omitempty" bson:"logo" binding:"required"` //  binding:"required"
	// Number of downloads
	Downloads int `json:"downloads,omitempty" form:"downloads,omitempty" query:"downloads,omitempty" bson:"downloads"`
	// Average score
	Score float64 `json:"score,omitempty" form:"score,omitempty" query:"score,omitempty" bson:"score"`
	// the development team
	Team string `json:"team,omitempty" form:"team,omitempty" query:"team,omitempty" bson:"team"`
	// the development teamIcon
	TeamIcon string `json:"teamIcon,omitempty" form:"teamIcon,omitempty" query:"teamIcon,omitempty" bson:"teamIcon"`
	// Template Properties
	Features string `json:"features,omitempty" form:"features,omitempty" query:"features,omitempty" bson:"features"`
	// preview
	Preview []system.SysBaseFiles `json:"preview,omitempty" form:"preview,omitempty" query:"preview,omitempty" bson:"preview"`
	// Copyright Notice
	Copyright string `json:"copyright,omitempty" form:"copyright,omitempty" query:"copyright,omitempty" bson:"copyright"`
	// development language
	DevLanguage []string `json:"devLanguage,omitempty" form:"devLanguage,omitempty" query:"devLanguage,omitempty" bson:"devLanguage"`
	// System Language
	SysLanguage []string `json:"sysLanguage,omitempty" form:"sysLanguage,omitempty" query:"sysLanguage,omitempty" bson:"sysLanguage"`
	// Template classification
	Category []string `json:"category,omitempty" form:"category,omitempty" query:"category,omitempty" bson:"category"`
	// poster
	Poster []system.SysBaseFiles `json:"poster,omitempty" form:"poster,omitempty" query:"poster,omitempty" bson:"poster"`
	// userID
	UserId string `json:"userId,omitempty" form:"userId,omitempty" query:"userId,omitempty" bson:"userId"`
	// Message notification
	Notice Notice `json:"notice,omitempty" form:"notice,omitempty" query:"notice,omitempty" bson:"notice"`
	// Software status 1. Software status 2. Software status 3. Software status 4. Software status
	State int `json:"state,omitempty" form:"state,omitempty" query:"state,omitempty" bson:"state"` //  binding:"required"
	// Delete or not
	IsDel bool `json:"isDel,omitempty" form:"isDel,omitempty" query:"isDel,omitempty" bson:"isDel"`
	// configuration information configuration information
	Config DocTemplateConfig `json:"config,omitempty" form:"config,omitempty" query:"config,omitempty" bson:"config"`
	// Software supported modes
	SupportMode Mode `json:"supportMode,omitempty" form:"supportMode,omitempty" query:"supportMode,omitempty" bson:"supportMode"`
	// Related software Related software
	RelationSoftIds []string `json:"relationSoftIds,omitempty" form:"relationSoftIds,omitempty" query:"relationSoftIds,omitempty" bson:"relationSoftIds"`
	// Related software
	RelationSoft []DocTemplate `json:"relationSoft,omitempty" form:"relationSoft,omitempty" query:"relationSoft,omitempty" bson:"relationSoft"`
}

// Notice Message notification
type Notice struct {
	Method string   `json:"method" form:"method" query:"method" bson:"method"` // Request Type Request TypeGET、POST
	Host   string   `json:"host" form:"host" query:"host" bson:"host"`         // ip/domain name
	Uri    string   `json:"uri" form:"uri" query:"uri" bson:"uri"`             // route
	Query  []string `json:"query" form:"query" query:"query" bson:"query"`     // parameter parameter version/details
}

// Mode Software supported modes
type Mode struct {
	// Dockermode
	Docker bool `json:"docker,omitempty" form:"docker,omitempty" query:"docker,omitempty" bson:"docker"`
	// Swarmmode
	Swarm bool `json:"swarm,omitempty" form:"swarm,omitempty" query:"swarm,omitempty" bson:"swarm"`
}
