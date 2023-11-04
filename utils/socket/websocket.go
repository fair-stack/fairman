// Package socket
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-25 10:20
package socket

import (
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Upgrade upgradegetupgradewebSocketupgrade
func Upgrade(c *gin.Context) (*websocket.Conn, error) {
	//upgradegetupgradewebSocketupgrade
	return upGrader.Upgrade(c.Writer, c.Request, nil)
	//if err != nil {
	//	response.FailWithMessage("websocketOpen failed！", c)
	//}
	//
}

// WsPing The only purpose of this method is to determinewebsocketThe only purpose of this method is to determine，The only purpose of this method is to determine，The only purpose of this method is to determine
func WsPing(ws *websocket.Conn) {
	// The only purpose of this method is to determinewebsocketThe only purpose of this method is to determine，The only purpose of this method is to determine，The only purpose of this method is to determine
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			return
		}
	}
}

func WsWriter(reader io.Reader, writer *websocket.Conn) {
	buf := make([]byte, 8192)
	for {
		nr, err := reader.Read(buf)
		if nr > 0 {
			fmt.Println("news： ", strings.Contains(string(buf[0:nr]), "vim"))
			if strings.Contains(string(buf[0:nr]), "exit") {
				writer.Close()
				//return
			} else if string(buf[0:nr]) == "" {
				continue
			}
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

func WsReader(reader *websocket.Conn, writer io.Writer) {
	for {
		messageType, p, err := reader.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(messageType)
		if messageType == websocket.TextMessage {
			if string(p) == "" {
				continue
			}
			fmt.Println("send： ", string(p))
			_, err = writer.Write(p)
			if err != nil {
				return
			}
		}
	}
}

// WsWriterLog WebSocket Get real-time logs
func WsWriterLog(ws *websocket.Conn, r io.ReadCloser) {
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
		//fmt.Println("Printing： ", string(dat))
		err = ws.WriteMessage(websocket.TextMessage, dat)
		if err != nil {
			return
		}
	}
}
