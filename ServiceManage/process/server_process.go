package process

import (
	"ServiceManage/common/redis"
	common_struct "ServiceManage/common/struct"
	"ServiceManage/common/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

var (
	wsServerToConn = make(map[string]map[string]*websocket.Conn)
	rwLock         = new(sync.RWMutex)
)

func RegisterRenewServer(serverName, serverAddress string, ws *websocket.Conn) {
	rwLock.Lock()
	defer rwLock.Unlock()
	err := redis.DB.RegisterRenewServer(serverName, serverAddress)
	if err != nil {
		fmt.Println("服务注册失败 err:" + err.Error())
		return
	}
	if m2, ok := wsServerToConn[serverName]; ok {
		m2[serverAddress] = ws
		wsServerToConn[serverName] = m2
	} else {
		m := make(map[string]*websocket.Conn)
		wsServerToConn[serverName] = m
	}
	dataSynchronizationSendMsg(serverName)
}

func DeleteServer(serverName, serverInfo string) {
	rwLock.Lock()
	defer rwLock.Unlock()
	redis.DB.DeleteServer(serverName, serverInfo)
	if m2, ok := wsServerToConn[serverName]; ok {
		if _, ok := m2[serverInfo]; ok {
			//ws.Close()
			delete(m2, serverInfo)
		}
	}
}

// SendMsg 返回服务信息,保证服务高可用的公共方法
func SendMsg(message common_struct.ServerProviderReturnMessage, conn *websocket.Conn) bool {
	bytes := utils.StructToJsonBytes(&message)
	err := conn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		fmt.Println("发送消息失败==" + utils.StructToJsonString(message))
		return false
	}
	return true
}

//服务状态更新后同步给全部服务
func dataSynchronizationSendMsg(serverName string) {
	if m, ok := wsServerToConn[serverName]; ok {
		message := common_struct.ServerProviderReturnMessage{CommonReturnMessage: struct {
			Code                              int64
			CentralDispatchServiceInformation []string
		}{Code: websocket.CloseProtocolError, CentralDispatchServiceInformation: []string{""}}}
		for _, ws := range m {
			SendMsg(message, ws)
		}
	}
}
func GetServerInfo(serverName string) map[string]*websocket.Conn {
	return wsServerToConn[serverName]
}
