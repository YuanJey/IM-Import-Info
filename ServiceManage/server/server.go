package server

import (
	"ServiceManage/auth"
	"ServiceManage/common/constant"
	common_struct "ServiceManage/common/struct"
	"ServiceManage/process"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func Register(c *gin.Context) {
	//鉴权
	if !auth.Authentication(c.GetHeader("token")) {
		return
	}
	//ip := c.ClientIP()
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	for {
		//读取ws中的数据
		//mt, message, err := ws.ReadMessage()
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}

		askMessage := common_struct.ServerProviderAskMessage{}
		_, err = processMessage(message, &askMessage)
		if err != nil {
			fmt.Println("数据错误")
			break
		}
		switch askMessage.Code {
		case constant.ServerRegister:
			process.RegisterRenewServer(askMessage.ServerName, askMessage.ProviderAddress, ws)
		case constant.ServerLogOut:
			process.DeleteServer(askMessage.ServerName, askMessage.ProviderAddress)
			ws.Close()
		case constant.Ping:
			process.RegisterRenewServer(askMessage.ServerName, askMessage.ProviderAddress, ws)
		}
	}
}
func processMessage(message []byte, AskMessage interface{}) (interface{}, error) {
	err := json.Unmarshal(message, &AskMessage)
	return AskMessage, err
}
