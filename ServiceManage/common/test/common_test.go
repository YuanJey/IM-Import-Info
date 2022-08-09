package db

import (
	"ServiceManage/common/config"
	"ServiceManage/common/redis"
	common_struct "ServiceManage/common/struct"
	"ServiceManage/common/utils"
	"context"
	"fmt"
	"testing"
)

func Test_Config(t *testing.T) {
	fmt.Println(config.Config.ClientType[0])
	fmt.Println(config.Config.Redis)
}
func Test_Redis(t *testing.T) {
	//redis.DB.SetSessionKey("asdasd", "asdasdasdas")
	redis.DB.RegisterRenewServer("asdasd", "asdasdasdas")

}

func Test_P(t *testing.T) {
	message := common_struct.ServerProviderAskMessage{
		CommonAskMessage: common_struct.CommonAskMessage{
			Code:       0,
			ServerName: "asdas",
		},
		ProviderAddress: "sadasd",
	}

	message1 := common_struct.ServerProviderAskMessage{}
	bytes := []byte(utils.StructToJsonString(message))
	utils.ByteToStruct(bytes, &message1)
	fmt.Println(utils.StructToJsonString(message1))
}

func Test_W(t *testing.T) {
	for {
		message, _ := redis.DB.Subscribe("aaa").ReceiveMessage(context.Background())
		fmt.Println(message)
	}
}
func Test_R(t *testing.T) {
	redis.DB.Publish("aaa", "asdasdasda")
}
