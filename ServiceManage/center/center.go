package center

import (
	"ServiceManage/auth"
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

func DataSynchronization(c *gin.Context) {
	//鉴权
	if auth.Authentication(c.GetHeader("token")) {
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
		fmt.Println(string(message))
	}

}
