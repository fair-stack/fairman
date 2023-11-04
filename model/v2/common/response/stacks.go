// Package response
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-22 15:03
package response

type Test struct {
	Name         string `json:"name" form:"name" query:"name" bson:"name"`
	SERVICES     string `json:"services" form:"services" query:"services" bson:"services"`
	ORCHESTRATOR string `json:"orchestrator" form:"orchestrator" query:"orchestrator" bson:"orchestrator"`
}
