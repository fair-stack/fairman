// Package response
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-11 17:17
package response

type RegistryResponse struct {
	Code int `json:"code"`
	//Data []docker.Registry `json:"data"`
	Msg string `json:"message"`
}
