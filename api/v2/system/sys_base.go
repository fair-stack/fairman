// Package system
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-03-16 17:13
package system

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/model/v2/common/response"
	"cnic/fairman-gin/model/v2/system"
	"cnic/fairman-gin/utils"
	"cnic/fairman-gin/utils/socket"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

//var upGrader = websocket.Upgrader{
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}
//
//// Login landing
//// @Summary User login
//// @Description User login
//// @Tags Docker Base
//// @Accept application/json
//// @Produce application/json
//// @Param object body request.UserRequest true "Login parameters"
//// @Success 200 {object} response.Response
//// @Router /v2/image/file [post]
//func (b *BaseApi) Login(c *gin.Context) {
//	var l *request.UserRequest
//
//	if errStr, err := utils.BaseValidator(&l, c); err != nil {
//		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
//		return
//	}
//	fmt.Println(l)
//	resp, err := baseService.Login(l)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	b.tokenNext(c, resp)
//}
//
//// Sign in after loginjwt
//func (b *BaseApi) tokenNext(c *gin.Context, user response.UserResponse) {
//	j := utils.NewJWT()
//	claims := j.CreateClaims(utils.BaseClaims{
//		UserID:   user.Id,
//		UserName: user.ZhName,
//	})
//	token, err := j.CreateToken(claims)
//	if err != nil {
//		global.FairLog.Error("obtaintokenobtain!", zap.Error(err))
//		response.FailWithMessage("obtaintokenobtain", c)
//		return
//	}
//
//	user.Token = token
//
//	response.OkWithData(user, c)
//}
//
//func (b *BaseApi) Restart(c *gin.Context) {
//
//	if err := baseService.Restart(); err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	response.Ok(c)
//}
//
//func (b *BaseApi) Test(c *gin.Context) {
//	//id := c.Query("id")
//	//test, s2, _ := softwareService.Test(id)
//	//
//	//response.OkWithDetailed(test, s2, c)
//	//response.OkWithData("4444", c)
//
//	// GetContainerLogsWs WebSocketGet container logs，Get container logs
//	// @Tags Container
//	// @Summary WebSocketGet container logs，Get container logs
//	// @Param containerId path string true "containerID"
//	// @Produce application/json
//	// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully obtained"}"
//	// @Router /v1/container/logs [get]
//	//containerId := c.Query("container")
//	//fmt.Println("containerId：", containerId)
//	fmt.Println("test")
//	//upgradegetupgradewebSocketupgrade
//	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		fmt.Println("upgrade:", err)
//		response.FailWithMessage("websocketOpen failed！", c)
//	}
//	defer func() {
//		_ = ws.Close()
//	}()
//
//	fmt.Println("test1")
//
//	url := "unix:///var/run/docker.sock"
//	//if runtime.GOOS == "windows" {
//	//	url = "npipe:////./pipe/docker_engine"
//	//}
//
//	cli, err := client.NewClientWithOpts(
//		client.FromEnv,
//		client.WithAPIVersionNegotiation(),
//		client.WithHost(url),
//	)
//	fmt.Println("cli：", err)
//
//	ctx := context.Background()
//	ir, err := cli.ContainerExecCreate(ctx, "1ef21ce51907", types.ExecConfig{
//		AttachStdin:  true,
//		AttachStdout: true,
//		AttachStderr: true,
//		Cmd:          []string{"sh"},
//		Tty:          true,
//	})
//
//	if err != nil {
//		return
//	}
//
//	hr, err := cli.ContainerExecAttach(ctx, ir.ID, types.ExecStartCheck{
//		Detach: false, Tty: true,
//	})
//	fmt.Println("r：", err)
//
//	defer hr.Close()
//	defer func() {
//		hr.Conn.Write([]byte("exit\r"))
//	}()
//
//	go func() {
//		WsWriter(hr.Conn, ws)
//	}()
//	WsReader(ws, hr.Conn)
//}
//
//func WsWriter(reader io.Reader, writer *websocket.Conn) {
//	buf := make([]byte, 8192)
//	for {
//		nr, err := reader.Read(buf)
//		if nr > 0 {
//			err := writer.WriteMessage(websocket.BinaryMessage, buf[0:nr])
//			if err != nil {
//				return
//			}
//		}
//		if err != nil {
//			return
//		}
//	}
//}
//
//func WsReader(reader *websocket.Conn, writer io.Writer) {
//	for {
//		messageType, p, err := reader.ReadMessage()
//		if err != nil {
//			return
//		}
//		if messageType == websocket.TextMessage {
//			fmt.Println(p)
//			_, err = writer.Write(p)
//			if err != nil {
//				return
//			}
//		}
//	}
//}

func (base *BaseApi) BaseLogs(c *gin.Context) {
	fmt.Println("logs####################################")
	//upgradegetupgradewebSocketupgrade
	ws, err := socket.Upgrade(c)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = ws.Close()
	}()

	global.BaseChan = make(chan global.ChanLogModal)

	go func() {
		for baseChan := range global.BaseChan {
			buf, err := baseChan.ToJsonByte()
			if err != nil {
				global.FairLog.Error(utils.Errorf(err).Error())
				continue
			}
			ws.WriteMessage(websocket.TextMessage, buf)
		}
	}()

	// readwsread
	func() {
		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				global.FairLog.Error(utils.Errorf(err).Error())
				break
			}

			fmt.Println(string(msg))

			if string(msg) == "close" {
				var respChan = global.ChanLogModal{
					Level:   "exit",
					Message: "The log service has been shut down, The log service has been shut down",
				}
				respBuf, _ := respChan.ToJsonByte()
				ws.WriteMessage(websocket.TextMessage, respBuf)
				break
			} else if string(msg) == "ping" {
				ws.WriteMessage(websocket.TextMessage, []byte("pong"))
			}
		}

		close(global.BaseChan)
	}()
}

// GetUnique Obtain unique identification uuid
// @Tags Base
// @Summary Obtain unique identification uuid
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully obtained"}"
// @Router /v1/base/uuid [get]
func (base *BaseApi) GetUnique(c *gin.Context) {
	unique, err := base.getUUID()
	if err != nil {
		response.FailWithMessage("Failed to obtain unique identity", c)
		return
	}
	response.OkWithData(unique, c)
}

func (base *BaseApi) getUUID() (string, error) {
	var basePath string
	if gin.Mode() == gin.DebugMode {
		basePath = ".fairman"
	} else {
		basePath = "/etc/.fairman"
	}

	f, err := os.Open(basePath + "/.unique")
	if err != nil {
		return "", err
	}

	b := make([]byte, 36)
	_, err = f.Read(b)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (base *BaseApi) UpdateApp(c *gin.Context) {
	var appVersion system.SysAppVersion

	if errStr, err := utils.BaseValidator(&appVersion, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	if err := baseService.UpdateApp(appVersion); err != nil {
		fmt.Println(err)
		response.FailWithMessage("Update failed", c)
		return
	}

	response.OkWithMessage("Update successful", c)
}

// UploadFile Upload files
// @Summary Upload files
// @Description Upload files
// @Tags Docker Base
// @Accept multipart/form-data
// @Produce application/json
// @Param file formData file true "file"
// @Success 200 {object} response.Response
// @Router /v2/base/file [post]
func (base *BaseApi) UploadFile(c *gin.Context) {
	fromData, _ := c.MultipartForm()
	files := fromData.File["file"]

	for _, file := range files {
		// TODO: Check if the file exists，Check if the file exists，Check if the file exists。
		extstring := path.Ext(file.Filename)
		//Generate a new file name based on the current timestamp
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt, 10)
		//New file name
		fileName := fmt.Sprintf("%s@%s", fileNameStr, file.Filename)
		//Save Upload File
		f, err := utils.Mkdir("upload")
		if err != nil {
			response.FailWithMessage("Path creation failed!", c)
			return
		}
		filePath := filepath.Join(f, "/", fileName)
		if utils.Exists(filePath) {
			response.FailCodeMessage(http.StatusBadRequest, "file already exist", c)
			return
		}
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			response.FailWithMessage("Failed to save file!", c)
			return
		}

		var baseFile = &system.SysBaseFiles{
			Name:   fileName,
			Url:    "/" + filePath,
			Size:   file.Size,
			Suffix: extstring,
		}

		//fileId, err := baseService.UploadFile(baseFile)
		//if err != nil {
		//	response.FailWithMessage("Failed to save file!", c)
		//	return
		//}
		//
		//baseFile.Id = fileId
		//baseFile.CreateAt = time.Now()
		response.OkWithData(baseFile, c)
	}
}
