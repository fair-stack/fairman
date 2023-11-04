// Package software
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-10 16:35
package software

type ServiceGroup struct {
	//SoftwareService
	//TemplateService
}

var (
	GetEndpointById     = "%s/v2/client/endpoint/%s"
	GetTemplateById     = "%s/v2/client/template/%s"
	GetRegistry         = "%s/v2/client/registry?template_id=%s"
	GetSoftwareById     = "%s/v2/client/software/%s"
	GetSoftwareByUserId = "%s/v2/client/software?user_id=%s"
	UpdSoftwareStatus   = "%s/v2/client/software/%s/status?status=%s"

	AddSoftware = "%s/v2/client/software"
)
