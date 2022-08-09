package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	etcd "go.etcd.io/etcd/client/v3"
)

var (
	client *etcd.Client
)

func init() {
	config := etcd.Config{
		Endpoints: []string{"127.0.0.1:2379"}}
	client, _ = etcd.New(config)
}
func main() {
	//RegisterRenewServer("asdasd", "asdasdasdas")
	TimedExpiration("asaaaad", "asdasdas", 20)
	Monitor("asaaaad")
}
func add(key, value string) error {
	_, err := client.Put(context.Background(), key, value)
	if err != nil {
		return err
	}
	return nil
}
func get(key string) error {
	_, err := client.Get(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}
func del(key string) error {
	_, err := client.Delete(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}

func TimedExpiration(key, value string, s int64) error {
	grant, err := client.Grant(context.Background(), s)
	if err != nil {
		return err
	}
	id := grant.ID
	_, err = client.Put(context.Background(), key, value, etcd.WithLease(id))
	if err != nil {
		return err
	}
	return nil
}

func Monitor(key string) {
	watchChan := client.Watch(context.Background(), key)
	for response := range watchChan {
		for _, ev := range response.Events {
			fmt.Printf("Type: %s Key:%s Value:%s  \n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			processing(ev)
		}
	}
}

func processing(ev *etcd.Event) {
	switch ev.Type {
	case mvccpb.Event_EventType(0):
		//switch string(ev.Kv.Key) {
		//case constant.UserServer:
		//
		//}
		fmt.Println("更新")
	case mvccpb.Event_EventType(1):
		fmt.Println("删除")
	}
}
func RegisterRenewServer(serverName, serverAddress string) error {
	return add(serverName, serverAddress)
}
