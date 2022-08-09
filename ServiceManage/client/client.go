package client

import (
	"ServiceManage/process"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetServerInfoReq struct {
	ServerName string `json:"serverName" binding:"required"`
}

func GetServerInfo(c *gin.Context) {
	params := GetServerInfoReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": process.GetServerInfo(params.ServerName)})
}
