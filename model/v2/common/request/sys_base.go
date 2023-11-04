// Package request
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-16 17:16
package request

type UserRequest struct {
	Account  string `json:"account" form:"account" query:"account" binding:"required"`
	Password string `json:"password" form:"password" query:"password" binding:"required"`
}
