// Package request
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 17:50
package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // Page number
	PageSize int `json:"pageSize" form:"pageSize"` // Size per page
}

// GetById Find by id structure
type GetById struct {
	ID string `json:"id" form:"id" uri:"id"` // ID
}

// GetByName Find by id structure
type GetByName struct {
	Name string `json:"name" form:"name" uri:"name"` // ID
}
