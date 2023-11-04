// Package response
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-30 20:39
package response

type Comment struct {
	CreateAt string  `json:"CreateAt,omitempty" form:"CreateAt,omitempty" query:"CreateAt,omitempty"`
	Comment  string  `json:"comment" form:"comment" query:"comment" bson:"comment"`         // Software review
	Score    float64 `json:"score" form:"score" query:"score" bson:"score"`                 // Software rating
	UserName string  `json:"user_name" form:"user_name" query:"user_name" bson:"user_name"` // User Name
}
