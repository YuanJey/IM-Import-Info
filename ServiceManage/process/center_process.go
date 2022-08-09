package process

import (
	"fmt"
	"net"
	"strings"
)

var (
	Manages []ManageCenter
	//governing = make(map[int]*[]ManageCenter)
	//governing = 0
	//P         *int
)

type ManageCenter struct {
	IP string
	//Index int
}

//func init() {
//	governing[0] = &Manages
//	P = &governing
//}
func init() {
	ip, err := GetOutBoundIP()
	if err != nil {
		fmt.Println("获取本机ip失败")
		return
	}
	center := ManageCenter{
		IP: ip,
		//Index: 0,
	}
	AddManageCenter(center)
}

//注册中心数据同步

//中心横向拓展，当增加时广播给全部服务提供方可服务消费方
//当中心无法访问，按级别切换中心

func AddManageCenter(center ManageCenter) {
	Manages = append(Manages, center)
}

func DeleteManageCenter(index int) {
	Manages = append(Manages[:index], Manages[index+1:]...)
	//*P = index
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
