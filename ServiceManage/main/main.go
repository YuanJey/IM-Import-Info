package main

import (
	"ServiceManage/center"
	"ServiceManage/client"
	"ServiceManage/server"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	serverRouterGroup := r.Group("/server")
	{
		serverRouterGroup.GET("/register", server.Register)
	}
	clientRouterGroup := r.Group("/client")
	{
		clientRouterGroup.POST("/get_serverInfo", client.GetServerInfo)
	}
	centerRouterGroup := r.Group("/registration_center")
	{
		centerRouterGroup.GET("/synchronize", center.DataSynchronization)
	}
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		fmt.Println("启动失败")
	}
}
