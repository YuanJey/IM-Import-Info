package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	_, b, _, _ = runtime.Caller(0)
	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../..")
)

var Config config

type config struct {
	Redis struct {
		DBAddress     []string `yaml:"dbAddress"`
		DBMaxIdle     int      `yaml:"dbMaxIdle"`
		DBMaxActive   int      `yaml:"dbMaxActive"`
		DBIdleTimeout int      `yaml:"dbIdleTimeout"`
		DBUserName    string   `yaml:"dbUserName"`
		DBPassWord    string   `yaml:"dbPassWord"`
		EnableCluster bool     `yaml:"enableCluster"`
	} `yaml:"redis"`

	ClientType []int32 `yaml:"clientType"`

	Provider struct {
		ServerExpiration int64 `yaml:"ServerExpiration"`
	} `yaml:"provider"`

	ETCD struct {
		Address []string `yaml:"address"`
	} `yaml:"etcd"`
}

func init() {
	//path, _ := os.Getwd()
	//bytes, err := ioutil.ReadFile(path + "/config/config.yaml")
	// if we cd Open-IM-Server/src/utils and run go test
	// it will panic cannot find config/config.yaml

	cfgName := os.Getenv("CONFIG_NAME")
	if len(cfgName) == 0 {
		cfgName = Root + "/config/config.yaml"
	}

	viper.SetConfigFile(cfgName)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	bytes, err := ioutil.ReadFile(cfgName)
	if err != nil {
		panic(err.Error())
	}
	if err = yaml.Unmarshal(bytes, &Config); err != nil {
		panic(err.Error())
	}
}
