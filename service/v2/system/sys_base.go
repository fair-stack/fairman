// Package system
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-16 17:17
package system

import (
	"bytes"
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/common/request"
	"cnic/fairman-gin/model/v2/common/response"
	"cnic/fairman-gin/model/v2/system"
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/common"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type BaseService struct {
}

func (b *BaseService) Login(l *request.UserRequest) (respUser response.UserResponse, err error) {
	var userRes response.LoginResponse
	// http://datamenu.casdc.cn/api/v1/findOrgByAccount
	//uri := fmt.Sprintf("%s/api/v1/findOrgByAccount", global.FairConf.System.Login)
	uri := "http://127.0.0.1:8099/v2/base/login"
	err = utils.Post(uri, l, "application/json", &userRes)
	if err != nil {
		fmt.Println(uri)
		return response.UserResponse{}, utils.Errorf(err)
	}

	fmt.Println(userRes)
	if userRes.Code == 0 {
		respUser.Id = userRes.Data.Id
		respUser.ZhName = userRes.Data.ZhName
		respUser.Token = userRes.Data.Token
		return respUser, nil
	} else {
		return respUser, utils.Errorf("Incorrect username or password。")
	}
}

func (b *BaseService) OldLogin(l *request.UserRequest) (respUser response.UserResponse, err error) {
	basePassword, err := base64.StdEncoding.DecodeString(l.Password)
	if err != nil {
		return response.UserResponse{}, utils.Errorf(err)
	}
	l.Password = string(basePassword)
	var userRes response.LoginResponse
	uri := fmt.Sprintf("%s/api/v1/findOrgByAccount", global.FairConf.System.Login)
	err = utils.Post(uri, l, "application/json", &userRes)
	fmt.Println(userRes)
	if err != nil {
		fmt.Println(uri)
		return response.UserResponse{}, utils.Errorf(err)
	}
	if userRes.Code == 200 {
		return respUser, nil
	} else {
		return respUser, utils.Errorf("returnCodereturn200, return。")
	}
}

// Restart Restart service
func (b *BaseService) Restart() error {

	// Get the latest version on the server
	// Determine if the versions are consistent
	// If inconsistent，If inconsistent
	// Obtain a new version of the binary package
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", fmt.Sprintf("kill -s HUP %s", strconv.Itoa(os.Getpid())))

	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		return utils.Errorf(err, stderr.String())
	}

	fmt.Println(string(output))
	// If consistent，If consistent

	return nil

}

// UpdateApp Update Service
func (b *BaseService) UpdateApp(appVersion system.SysAppVersion) (err error) {
	// /mnt/data/fairman/web/dist
	// /mnt/data/fairman/api/fairman-gin
	// 1. Back up front-end and back-end files
	// 1.1 Backend file backup
	var cmd *exec.Cmd
	var stderr bytes.Buffer

	global.BaseChan <- global.ChanLogModal{
		Message:  "Start preparing to update the environment...",
		Level:    "success",
		Progress: 0,
	}
	time.Sleep(time.Millisecond * 500)
	// Backend Path
	backendPath := fmt.Sprintf("%s/fairman-gin", global.FairConf.System.BackendPath)
	// Front end path
	frontendPath := fmt.Sprintf("%s/dist", global.FairConf.System.FrontendPath)
	// Path for backend backup
	backendBakPath := fmt.Sprintf("%s/fairman-gin-%s-%d", global.FairConf.System.BackendPath, appVersion.Version, time.Now().Unix())
	// Path for front-end backup
	frontendBakPath := fmt.Sprintf("%s/dist-%s-%d", global.FairConf.System.FrontendPath, appVersion.Version, time.Now().Unix())

	global.BaseChan <- global.ChanLogModal{
		Message:  "Update environment preparation completed...",
		Level:    "success",
		Progress: 0.1,
	}

	fmt.Println("Backend Path：", backendPath)
	fmt.Println("Front end path：", frontendPath)
	fmt.Println("Path for backend backup：", backendBakPath)
	fmt.Println("Path for front-end backup：", frontendBakPath)
	// Restore backup

	global.BaseChan <- global.ChanLogModal{
		Message:  "Start backing up files...",
		Level:    "success",
		Progress: 0.1,
	}
	time.Sleep(time.Millisecond * 500)

	// Check if the file exists
	// 1.1 Backend file backup
	if ok, err := utils.PathExists(backendPath); ok {
		fmt.Println(utils.Errorf(err))
		err := common.MV(backendPath, backendBakPath)
		if err != nil {
			// 1.1.1 mvBackend file error，Backend file error，Backend file error
			global.BaseChan <- global.ChanLogModal{
				Message:  "Backend file backup failed，Backend file backup failed...",
				Level:    "error",
				Progress: 1,
			}
			return utils.Errorf(err)
		}
		fmt.Println("Backend backup successful")
	}

	// 1.2 Frontend file backup
	if ok, err := utils.PathExists(frontendPath); ok {
		fmt.Println(utils.Errorf(err))

		err := common.MV(frontendPath, frontendBakPath)
		if err != nil {
			// 1.2.1 Front end file backup failed，Front end file backup failed
			err := common.MV(backendBakPath, backendPath)
			fmt.Println(utils.Errorf(err))
			global.BaseChan <- global.ChanLogModal{
				Message:  "Front end file backup failed，Front end file backup failed，Front end file backup failed...",
				Level:    "error",
				Progress: 1,
			}
			return utils.Errorf(err)
		}
		fmt.Println("Front end backup successful")
	}

	global.BaseChan <- global.ChanLogModal{
		Message:  "File backup completed...",
		Level:    "success",
		Progress: 0.2,
	}
	global.BaseChan <- global.ChanLogModal{
		Message:  "Start downloading front-end files...",
		Level:    "success",
		Progress: 0.2,
	}
	time.Sleep(time.Millisecond * 500)

	// 2. Download the new front-end and back-end files to the specified directory
	// 2.1 Download front-end files
	err = utils.DownLoadFile(
		fmt.Sprintf("%s%s", global.FairConf.Service.Url, appVersion.Frontend.Url),
		fmt.Sprintf("%s/dist.zip", global.FairConf.System.FrontendPath),
	)
	if err != nil {
		// Failed to download file front-end file，Failed to download file front-end file
		fmt.Println(utils.Errorf(err))
		global.BaseChan <- global.ChanLogModal{
			Message:  "Front end file download failed，Front end file download failed，Front end file download failed...",
			Level:    "error",
			Progress: 1,
		}
		goto OnRevert
	}
	fmt.Println("Front end download successful")
	global.BaseChan <- global.ChanLogModal{
		Message:  "Successfully downloaded front-end files...",
		Level:    "success",
		Progress: 0.4,
	}
	global.BaseChan <- global.ChanLogModal{
		Message:  "Start downloading backend files...",
		Level:    "success",
		Progress: 0.4,
	}
	time.Sleep(time.Millisecond * 500)

	// Try deleting it，Try deleting it，Try deleting it
	os.Remove(fmt.Sprintf("%s/fairman-gin", global.FairConf.System.BackendPath))

	fmt.Println("Successfully deleted backend file")

	// 2.2 Download backend files
	err = utils.DownLoadFile(
		fmt.Sprintf("%s%s", global.FairConf.Service.Url, appVersion.Backend.Url),
		fmt.Sprintf("%s/fairman-gin", global.FairConf.System.BackendPath),
	)
	if err != nil {
		fmt.Println(utils.Errorf(err))
		global.BaseChan <- global.ChanLogModal{
			Message:  "Backend file download failed，Backend file download failed，Backend file download failed...",
			Level:    "error",
			Progress: 1,
		}
		goto OnRevert
	}

	fmt.Println("Backend download successful")
	global.BaseChan <- global.ChanLogModal{
		Message:  "Backend file download successful...",
		Level:    "success",
		Progress: 0.6,
	}
	global.BaseChan <- global.ChanLogModal{
		Message:  "Start restarting service task...",
		Level:    "success",
		Progress: 0.6,
	}
	time.Sleep(time.Second * 2)

	err = exec.Command("chmod", "+x", fmt.Sprintf("%s/fairman-gin", global.FairConf.System.BackendPath)).Run()
	if err != nil {
		return err
	}
	fmt.Println("Successfully modified backend file permissions")
	global.BaseChan <- global.ChanLogModal{
		Message:  "Modify file permissions...",
		Level:    "success",
		Progress: 0.7,
	}
	// 3. Restart service
	// 3.1 Decompress front-end files
	err = utils.Unzip(
		fmt.Sprintf("%s/dist.zip", global.FairConf.System.FrontendPath),
		fmt.Sprintf(global.FairConf.System.FrontendPath),
	)
	if err != nil {
		fmt.Println(utils.Errorf(err))
		global.BaseChan <- global.ChanLogModal{
			Message:  "Failed to extract file，Failed to extract file，Failed to extract file...",
			Level:    "error",
			Progress: 1,
		}
		goto OnRevert
	}
	fmt.Println("Front end decompression successful")
	global.BaseChan <- global.ChanLogModal{
		Message:  "Decompress files...",
		Level:    "success",
		Progress: 0.8,
	}
	time.Sleep(time.Millisecond * 500)

	// 3.2 Delete compressed package
	err = os.Remove(fmt.Sprintf("%s/dist.zip", global.FairConf.System.FrontendPath))
	if err != nil {
		return utils.Errorf(err)
	}
	fmt.Println("Successfully deleted front-end compressed package")

	fmt.Println(os.Getpid())

	//global.FairStore.Close()
	//global.FairDB.Close()

	global.BaseChan <- global.ChanLogModal{
		Message:  "Clean up old files...",
		Level:    "success",
		Progress: 0.9,
	}
	global.BaseChan <- global.ChanLogModal{
		Message:  "Clean up old files...",
		Level:    "success",
		Progress: 0.9,
	}
	time.Sleep(time.Second * 2)

	// 3.3 Restart service
	cmd = exec.Command("bash", "-c", fmt.Sprintf("kill -s HUP %s", strconv.Itoa(os.Getpid())))

	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(utils.Errorf(err))
		global.BaseChan <- global.ChanLogModal{
			Message:  "Restart service failed，Restart service failed，Restart service failed...",
			Level:    "error",
			Progress: 1,
		}
		goto OnRevert
	}
	fmt.Println("Service restarted successfully")
	global.BaseChan <- global.ChanLogModal{
		Message:  "Update completed，Update completed...",
		Level:    "success",
		Progress: 1,
	}
	time.Sleep(time.Second * 5)

	//fmt.Println(string(output))
	return nil

OnRevert:
	// 4. Restore Environment
	// 4.1 Restore front-end files
	common.MV(frontendBakPath, frontendPath)
	// 4.2 Restore backend files
	common.MV(backendBakPath, backendPath)
	// 4.3 Restart service
	cmd = exec.Command("bash", "-c", fmt.Sprintf("kill -s HUP %s", strconv.Itoa(os.Getpid())))
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(utils.Errorf(err))
		global.BaseChan <- global.ChanLogModal{
			Message:  "Restart service failed，Restart service failed...",
			Level:    "error",
			Progress: 1,
		}
		return utils.Errorf("Restart service failed")
	}

	fmt.Println(utils.Errorf(err))
	global.BaseChan <- global.ChanLogModal{
		Message:  "Update failed，Update failed...",
		Level:    "error",
		Progress: 1,
	}
	return utils.Errorf("Update failed")
}

var ifChannelsMapInit = false
var channelsMap = map[string]chan string{}

func (b *BaseService) initChannelsMap() {
	channelsMap = make(map[string]chan string)
}

func (b *BaseService) addChannel(traceId string) {
	if !ifChannelsMapInit {
		b.initChannelsMap()
		ifChannelsMapInit = true
	}

	var newChannel = make(chan string)

	channelsMap[traceId] = newChannel

	fmt.Println("Build SSE connection for user = " + traceId)
}

// BaseSEE service
func (b *BaseService) BaseSEE(traceId string, c *gin.Context) {
	//v, ok := c.Get("clientChan")
	//if !ok {
	//	return
	//}
	//clientChan, ok := v.(global.ClientChan)
	//if !ok {
	//	return
	//}

	//fmt.Println("traceId1: ", traceId)
	//c.Stream(func(w io.Writer) bool {
	//	// Stream message to client from message channel
	//	if msg, ok := <-global.Test[traceId]; ok {
	//		str, _ := msg.ToJsonStr()
	//		c.SSEvent("message", str)
	//		return true
	//	}
	//	return false
	//})
	fmt.Println(222)
}

type T1 struct {
	TraceId string `json:"trace_id,omitempty"`
	Message string `json:"message,omitempty"`
}

func SendNotification(traceId string, message string) {
	//log.Infof("Send notification to user = " + userEmail)
	//var msg = dao.NotificationLog{}
	//msg.MessageBody = messageBody
	//msg.UserEmail = userEmail
	//msg.Type = actionType
	//msg.Status = UNREAD
	//msg.CreateTime = time.Now()
	//msg.Create()
	var t1 T1
	t1.TraceId = traceId
	t1.Message = message
	msgBytes, _ := json.Marshal(t1)
	for key := range channelsMap {
		if strings.Contains(key, traceId) {
			channel := channelsMap[key]
			channel <- string(msgBytes)
		}
	}
}
