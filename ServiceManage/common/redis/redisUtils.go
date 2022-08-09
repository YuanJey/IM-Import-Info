package redis

import (
	"ServiceManage/common/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"reflect"
	"time"
)

const (
	accountTempCode       = "ACCOUNT_TEMP_CODE"
	resetPwdTempCode      = "RESET_PWD_TEMP_CODE"
	userIncrSeq           = "REDIS_USER_INCR_SEQ:" // user incr seq
	CurrentMaximumUserId  = "CURRENT_MAXIMUM_USERID"
	InitializingAnAccount = "9999999"
	ENCRYPTION            = "ENCRYPTION:"

	ServerProviderInfo = "SERVERPROVIDERINFO:"
)

func (d *DataBases) JudgeAccountEXISTS(account string) (bool, error) {
	key := accountTempCode + account
	n, err := d.rdb.Exists(context.Background(), key).Result()
	if n > 0 {
		return true, err
	} else {
		return false, err
	}
}
func (d *DataBases) SetAccountCode(account string, code, ttl int) (err error) {
	key := accountTempCode + account
	return d.rdb.Set(context.Background(), key, code, time.Duration(ttl)*time.Second).Err()
}
func (d *DataBases) GetAccountCode(account string) (string, error) {
	key := accountTempCode + account
	return d.rdb.Get(context.Background(), key).Result()
}

//Perform seq auto-increment operation of user messages
func (d *DataBases) IncrUserSeq(uid string) (uint64, error) {
	key := userIncrSeq + uid
	seq, err := d.rdb.Incr(context.Background(), key).Result()
	return uint64(seq), err
}

func (d *DataBases) GetMobile(key string) (bool, error) {
	key = "PUSH:" + key
	exec, err := d.rdb.Get(context.Background(), key).Result()
	if exec == "" {
		return false, err
	}
	return true, err
}

func (d *DataBases) AddMobile(key, PlatFormID string) error {
	key = "PUSH:" + key
	_, err := d.rdb.Set(context.Background(), key, PlatFormID, 0).Result()
	return err
}

func (d *DataBases) DelMobile(key string) error {
	key = "PUSH:" + key
	_, err := d.rdb.Del(context.Background(), key).Result()
	return err
}

func (d *DataBases) GeneratingUserID() (string, error) {
	_, err2 := d.rdb.Do(context.Background(), "GET", CurrentMaximumUserId).Result()
	if err2 != nil {
		d.rdb.Do(context.Background(), "SET", CurrentMaximumUserId, InitializingAnAccount)
	}
	id, err := d.rdb.Do(context.Background(), "INCR", CurrentMaximumUserId).Result()
	if err != nil {
		return "", nil
	}
	return cast.ToString(id), nil
}

func (d *DataBases) GeneratingUserID2() (string, error) {
	//_, err2 := d.rdb.Do(context.Background(), "GET", CurrentMaximumUserId).Result()
	_, err2 := d.rdb.Get(context.Background(), CurrentMaximumUserId).Result()
	if err2 != nil {
		//d.rdb.Do(context.Background(), "SET", CurrentMaximumUserId, InitializingAnAccount)
		d.rdb.Set(context.Background(), CurrentMaximumUserId, InitializingAnAccount, 0)
	}
	id, err := d.rdb.Incr(context.Background(), CurrentMaximumUserId).Result()
	if err != nil {
		return "", nil
	}
	return cast.ToString(id), nil
}

func (d *DataBases) GetSessionKey(key string) (string, error) {
	key = ENCRYPTION + key
	result, err := d.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (d *DataBases) SetSessionKey(key, SessionKey string) error {
	key = ENCRYPTION + key
	_, err := d.rdb.Set(context.Background(), key, SessionKey, 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func (d *DataBases) CommonClearCache(keys string) {
	result, err := d.rdb.Do(context.Background(), "keys", keys+"*").Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	if reflect.TypeOf(result).Kind() == reflect.Slice {
		val := reflect.ValueOf(result)
		if val.Len() == 0 {
			return
		}
		for i := 0; i < val.Len(); i++ {
			d.rdb.Del(context.Background(), val.Index(i).Interface().(string))
			//fmt.Printf("删除了rediskey：：%s \n", val.Index(i).Interface().(string))
		}
	}
}

// RegisterRenewServer 服务过期重新注册和续约
func (d *DataBases) RegisterRenewServer(serverName, serverAddress string) error {
	key := ServerProviderInfo + serverName + ":" + serverAddress
	_, err := d.rdb.Set(context.Background(), key, serverAddress, 0).Result()
	if err != nil {
		return err
	}
	d.rdb.ExpireAt(context.Background(), key, time.Now().Add(time.Duration(config.Config.Provider.ServerExpiration)*time.Second)).Result()
	return nil
}

func (d *DataBases) DeleteServer(serverName, serverAddress string) error {
	key := ServerProviderInfo + serverName + ":" + serverAddress
	_, err := d.rdb.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (d *DataBases) Publish(channel, message string) error {
	_, err := d.rdb.Publish(context.Background(), channel, message).Result()
	return err
}
func (d *DataBases) Subscribe(channel string) *redis.PubSub {
	return d.rdb.Subscribe(context.Background(), channel)
}
