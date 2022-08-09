package redis

import (
	"ServiceManage/common/config"
	"context"
	"fmt"
	go_redis "github.com/go-redis/redis/v8"
	"time"
)

var DB DataBases

type DataBases struct {
	rdb go_redis.UniversalClient
}

func init() {
	fmt.Println("tes", config.Config.Redis.DBUserName, config.Config.Redis.DBPassWord)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if config.Config.Redis.EnableCluster {
		DB.rdb = go_redis.NewClusterClient(&go_redis.ClusterOptions{
			Addrs:    config.Config.Redis.DBAddress,
			Username: config.Config.Redis.DBUserName,
			Password: config.Config.Redis.DBPassWord, // no password set
			PoolSize: 50,
		})
		_, err := DB.rdb.Ping(ctx).Result()
		if err != nil {
			panic(err.Error())
		}
	} else {
		DB.rdb = go_redis.NewClient(&go_redis.Options{
			Addr:     config.Config.Redis.DBAddress[0],
			Username: config.Config.Redis.DBUserName,
			Password: config.Config.Redis.DBPassWord, // no password set
			DB:       1,                              // use default DB
			PoolSize: 100,                            // 连接池大小
		})
		_, err := DB.rdb.Ping(ctx).Result()
		if err != nil {
			panic(err.Error())
		}
	}
}
