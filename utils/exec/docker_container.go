// Package exec
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-04-14 10:39
package exec

import (
	"bytes"
	"cnic/fairman-gin/global"
	dockerV2 "cnic/fairman-gin/model/v2/docker"
	"cnic/fairman-gin/utils"
	"context"
	"encoding/binary"
	"io"
	"time"

	"github.com/docker/docker/client"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gorilla/websocket"
)

type ContainerOperation struct {
	Endpoint dockerV2.DocEndpoint
	Cli      *client.Client
}

//// TODO: test??????????
//// GetContainerList Get a list of all containers（Get a list of all containers）
//func (c *ContainerOperation) GetContainerList() ([]types.Container, error) {
//	arg := filters.NewArgs(filters.KeyValuePair{Key: "label", Value: "com.docker.compose.project=test1"})
//
//	dockerClient, err := global.FairClient.CreateClient(c.Endpoint)
//	if err != nil {
//		return nil, utils.Errorf(err, "get container list failed")
//	}
//	defer dockerClient.Close()
//
//	return dockerClient.ContainerList(context.Background(), types.ContainerListOptions{All: true, Filters: arg})
//	// TODO: Need to remove software from non software stores，Need to remove software from non software stores
//}

// OperatorContainer Operation container
func (c *ContainerOperation) OperatorContainer(containerId string, operator string) error {
	ctx := context.Background()

	dockerClient, err := global.FairClient.CreateClient(c.Endpoint)
	if err != nil {
		return utils.Errorf(err, "operator container error")
	}
	defer dockerClient.Close()

	switch operator {
	case "start": // start-up
		options := types.ContainerStartOptions{}
		return dockerClient.ContainerStart(ctx, containerId, options)
	case "stop": // cease
		return dockerClient.ContainerStop(ctx, containerId, nil)
	case "restart": // restart
		duration := time.Second * 20
		return dockerClient.ContainerRestart(ctx, containerId, &duration)
	case "pause": // suspend
		return dockerClient.ContainerPause(ctx, containerId)
	case "unpause": // recovery
		return dockerClient.ContainerUnpause(ctx, containerId)
	case "prune": // clean up
		var filter filters.Args
		_, err := dockerClient.ContainersPrune(ctx, filter)
		return err
	default:
		return nil
	}
}

// RenameContainer Rename Container
func (c *ContainerOperation) RenameContainer(containerId string, newName string) error {
	dockerClient, err := global.FairClient.CreateClient(c.Endpoint)
	if err != nil {
		return utils.Errorf(err, "rename container error")
	}
	defer dockerClient.Close()

	return dockerClient.ContainerRename(context.Background(), containerId, newName)
}

// RemoveContainer Delete Container
func (c *ContainerOperation) RemoveContainer(containerId string, opt types.ContainerRemoveOptions) error {
	dockerClient, err := global.FairClient.CreateClient(c.Endpoint)
	if err != nil {
		return utils.Errorf(err, "remove container error")
	}
	defer dockerClient.Close()

	return dockerClient.ContainerRemove(context.Background(), containerId, opt)
}

// GetContainerInfo Container information
func (c *ContainerOperation) GetContainerInfo(containerIdOrName string) (types.ContainerJSON, error) {
	dockerClient, err := global.FairClient.CreateClient(c.Endpoint)
	if err != nil {
		return types.ContainerJSON{}, utils.Errorf(err, "get container info error")
	}
	defer dockerClient.Close()

	return dockerClient.ContainerInspect(context.Background(), containerIdOrName)
}

// GetContainerLogs Obtain 200 rows of logs
func (c *ContainerOperation) GetContainerLogs(containerId string, tail string) (string, error) {
	options := types.ContainerLogsOptions{ShowStdout: true, Timestamps: true}

	if tail != "" {
		options.Tail = tail
	}

	dockerClient, err := global.FairClient.CreateClient(c.Endpoint)
	if err != nil {
		return "", utils.Errorf(err, "get container logs error")
	}
	defer dockerClient.Close()

	logs, err := dockerClient.ContainerLogs(context.Background(), containerId, options)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(logs)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// GetContainerLogsWs WebSocket Get real-time logs
func (c *ContainerOperation) GetContainerLogsWs(containerId string) (io.ReadCloser, error) {
	dockerClient, err := global.FairClient.CreateClient(c.Endpoint)
	if err != nil {
		return nil, utils.Errorf(err, "get container logs error")
	}
	defer dockerClient.Close()

	return dockerClient.ContainerLogs(context.Background(), containerId, types.ContainerLogsOptions{
		ShowStderr: true,
		ShowStdout: true,
		Follow:     true,
		//Tail:       "40",
	})
}

// WsWriterLog WebSocket Get real-time logs
func (c *ContainerOperation) WsWriterLog(ws *websocket.Conn, r io.ReadCloser) {
	hdr := make([]byte, 8)
	for {
		_, err := r.Read(hdr)
		if err != nil || err == io.EOF {
			return
		}
		//var w io.Writer
		//fmt.Println(string(hdr))
		//switch hdr[0] {
		//case 1:
		//	w = os.Stdout
		//default:
		//	w = os.Stderr
		//}
		count := binary.BigEndian.Uint32(hdr[4:])
		dat := make([]byte, count)
		_, err = r.Read(dat)
		if err != nil {
			return
		}
		//fmt.Println("Printing： ", dat)
		err = ws.WriteMessage(websocket.BinaryMessage, dat)
		if err != nil {
			return
		}
	}
}

// WsPing The only purpose of this method is to determinewebsocketThe only purpose of this method is to determine，The only purpose of this method is to determine，The only purpose of this method is to determine
func (c *ContainerOperation) WsPing(ws *websocket.Conn) {
	// The only purpose of this method is to determinewebsocketThe only purpose of this method is to determine，The only purpose of this method is to determine，The only purpose of this method is to determine
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			return
		}
	}
}

// ExportContainer Export Container
func (c *ContainerOperation) ExportContainer(containerId string) (io.ReadCloser, error) {
	dockerClient, err := global.FairClient.CreateClient(c.Endpoint)
	if err != nil {
		return nil, utils.Errorf(err, "export container error")
	}
	defer dockerClient.Close()

	return dockerClient.ContainerExport(context.Background(), containerId)
}

// MonitorContainer monitorDocker
func (c *ContainerOperation) MonitorContainer(containerId string) (types.ContainerStats, error) {
	dockerClient, err := global.FairClient.CreateClient(c.Endpoint)
	if err != nil {
		return types.ContainerStats{}, utils.Errorf(err, "monitor container error")
	}
	defer dockerClient.Close()

	return dockerClient.ContainerStats(context.Background(), containerId, false)
}

// WsWriterMonitor WebSocket Real time monitoringDocker
func (c *ContainerOperation) WsWriterMonitor(ws *websocket.Conn, containerId string) {
	for {
		containerStats, err := c.MonitorContainer(containerId)
		buf := new(bytes.Buffer)
		//io.ReadCloser convert to Buffer convert tojsonconvert to
		_, _ = buf.ReadFrom(containerStats.Body)
		newStr := buf.Bytes()
		err = ws.WriteMessage(websocket.BinaryMessage, newStr)
		if err != nil {
			return
		}
		time.Sleep(time.Second * 10)
	}
}

// GetImageExist adoptImageHistoryadopt，adopt
//func (c *ContainerOperation) GetImageExist(imageName string) error {
//	global.FMC_CHAL <- fmt.Sprintf("inspect[%s]inspect", imageName)
//	_, err := dockerClient.ImageHistory(context.Background(), imageName)
//
//	if err != nil {
//		global.FMC_CHAL <- fmt.Sprintf("[%s]Not downloaded，Not downloaded...", imageName)
//		for i := 0; i < 3; i++ {
//			reader, err := dockerClient.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
//			if err != nil {
//				global.FMC_LOG.Error(err.Error())
//				global.FMC_CHAL <- fmt.Sprintf("Section[%d]Section Section[%s]Section，Section，Section：(%s)", i, imageName, err.Error())
//				//return errors.New("Failed to obtain software, Failed to obtain software")
//			}
//			//buf := new(bytes.Buffer)
//			scanner := bufio.NewScanner(reader)
//			const maxCapacity = 512 * 1024
//			buf := make([]byte, maxCapacity)
//			scanner.Buffer(buf, maxCapacity)
//			fmt.Println("Start reading download logs")
//			global.FMC_CHAL <- fmt.Sprintf("[%s]Start downloading...", imageName)
//			for scanner.Scan() {
//				global.FMC_CHAL <- scanner.Text()
//			}
//			if err != nil {
//				global.FMC_CHAL <- fmt.Sprintf("Section[%d]Section Section[%s]Section，Section，Section：(%s)", i, imageName, err.Error())
//				//return err
//			} else {
//				global.FMC_CHAL <- fmt.Sprintf("[%s]Download completed，Download completed...", imageName)
//				return nil
//			}
//		}
//	}
//	global.FMC_CHAL <- fmt.Sprintf("[%s]Downloaded，Downloaded...", imageName)
//	return err
//}

// ContainerTerminal WebSocket Real time terminal
func (c *ContainerOperation) ContainerTerminal(containerId string, cmd string) (hr types.HijackedResponse, err error) {
	ctx := context.Background()

	dockerClient, err := global.FairClient.CreateClient(c.Endpoint)
	if err != nil {
		return types.HijackedResponse{}, utils.Errorf(err, "websocket container terminal")
	}
	defer dockerClient.Close()

	ir, err := dockerClient.ContainerExecCreate(ctx, containerId, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{cmd},
		Tty:          true,
	})

	if err != nil {
		return
	}

	hr, err = dockerClient.ContainerExecAttach(ctx, ir.ID, types.ExecStartCheck{
		Detach: false, Tty: true,
	})

	if err != nil {
		return
	}
	return
}

func (c *ContainerOperation) WsWriter(reader io.Reader, writer *websocket.Conn) {
	buf := make([]byte, 8192)
	for {
		nr, err := reader.Read(buf)
		if nr > 0 {
			err := writer.WriteMessage(websocket.BinaryMessage, buf[0:nr])
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
	}
}

func (c *ContainerOperation) WsReader(reader *websocket.Conn, writer io.Writer) {
	for {
		messageType, p, err := reader.ReadMessage()
		if err != nil {
			return
		}
		if messageType == websocket.TextMessage {
			_, err = writer.Write(p)
			if err != nil {
				return
			}
		}
	}
}
