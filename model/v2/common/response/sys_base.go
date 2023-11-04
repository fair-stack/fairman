// Package response
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-16 17:20
package response

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id          string `json:"id"`
		ZhName      string `json:"zh_Name"`
		Token       string `json:"token"`
		RoleType    string `json:"roleType"`
		AuthorityId string `json:"authority_id"`
	} `json:"data"`
}

type OldLoginResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		// userid(user)
		Id string `json:"id,omitempty" form:"id,omitempty" query:"id,omitempty"`
		// user name(user name)
		Username string `json:"username" form:"username" query:"username" bson:"username"`
		// dense  dense(dense)
		Password string `json:"password" form:"password" query:"password" bson:"password"`
		// Chinese name(Chinese name)
		ZHName string `json:"zh_name" form:"zh_name" query:"zh_name" bson:"zh_name"`
		// type(type)
		Type string `json:"type" form:"type" query:"type" bson:"type"`
		// state(state)
		Status string `json:"status" form:"status" query:"status" bson:"status"`
	} `json:"result"`
}

type ClientUserResponse struct {
}

type UserResponse struct {
	Id     string `json:"id" mapstructure:"id"`
	ZhName string `json:"zh_Name" mapstructure:"zh_Name"`
	Token  string `json:"token"`
}
