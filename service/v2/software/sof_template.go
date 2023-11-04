// Package software
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-28 10:59
package software

//import (
//	"cnic/fairman-gin/global"
//	"cnic/fairman-gin/model/v2/common/response"
//	"cnic/fairman-gin/model/v2/docker"
//	"cnic/fairman-gin/utils"
//	"fmt"
//
//	"github.com/docker/cli/cli/compose/loader"
//	cliTypes "github.com/docker/cli/cli/compose/types"
//)
//
//type TemplateService struct {
//}
//
//func (t *TemplateService) GetTemplateConfigById(id string) (map[string]response.TemplateConfigResp, error) {
//	var resp response.Response
//	var template docker.CustomTemplate
//	templateResp := make(map[string]response.TemplateConfigResp)
//	url := fmt.Sprintf(GetTemplateById, global.FairConf.Service.Url, id)
//	err := utils.Get(url, &resp)
//	if err != nil {
//		return nil, err
//	}
//
//	if resp.Code == 0 {
//		err1 := utils.DataJson(resp.Data, &template)
//		if err1 != nil {
//			return nil, err1
//		}
//
//		composeConfigYAML, err := loader.ParseYAML([]byte(template.FileContent))
//		if err != nil {
//			return nil, utils.Errorf("The software configuration is abnormal，The software configuration is abnormal")
//		}
//
//		composeConfigFile := cliTypes.ConfigFile{
//			Config: composeConfigYAML,
//		}
//
//		composeConfigDetails := cliTypes.ConfigDetails{
//			ConfigFiles: []cliTypes.ConfigFile{composeConfigFile},
//			Environment: map[string]string{},
//		}
//
//		composeConfig, err := loader.Load(composeConfigDetails, func(options *loader.Options) {
//			options.SkipValidation = true
//			options.SkipInterpolation = true
//		})
//		if err != nil {
//			return nil, utils.Errorf(err)
//		}
//
//		for key := range composeConfig.Services {
//			service := composeConfig.Services[key]
//			templateResp[service.Name] = response.TemplateConfigResp{
//				Ports:       service.Ports,
//				Environment: service.Environment,
//				Mounts:      service.Volumes,
//				Hostname:    service.Hostname,
//			}
//		}
//
//		return templateResp, nil
//	}
//	return nil, utils.Errorf("code != 0")
//}
//
//func (t *TemplateService) GetTemplateById(id string) (docker.CustomTemplate, error) {
//	var resp response.Response
//	var template docker.CustomTemplate
//	url := fmt.Sprintf(GetTemplateById, global.FairConf.Service.Url, id)
//	err := utils.Get(url, &resp)
//	if err != nil {
//		return template, err
//	}
//
//	if resp.Code == 0 {
//		err1 := utils.DataJson(resp.Data, &template)
//		if err1 != nil {
//			return template, err1
//		}
//
//		return template, nil
//	}
//	return template, utils.Errorf("code != 0")
//}
