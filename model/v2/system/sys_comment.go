// Package system
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-30 20:43
package system

type Comment struct {
	Comment    string  `json:"comment" form:"comment" query:"comment" bson:"comment"`                 // Software review
	Score      float64 `json:"score" form:"score" query:"score" bson:"score"`                         // Software rating
	TemplateId string  `json:"template_id" form:"template_id" query:"template_id" bson:"template_id"` // softwareID
	UserId     string  `json:"user_id" form:"user_id" query:"user_id" bson:"user_id"`                 // userID
	UserName   string  `json:"user_name" form:"user_name" query:"user_name" bson:"user_name"`         // User Name
}
