package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

var Cfg *ConfigMap

type ConfigMap struct {
	ListenPort string `json:"ListenPort" yaml:"listen_port"`

	LogPath string `json:"LogPath" yaml:"log_path"`

	RedisPassword  string `json:"RedisPassword" yaml:"redis_password"`
	RedisIsCluster bool   `json:"RedisIsCluster,string" yaml:"redis_is_cluster"`

	RedisAddr         string   `json:"RedisAddr" yaml:"redis_addr"`
	RedisClusterNodes []string `json:"RedisClusterNodes" yaml:"redis_cluster_nodes"`

	MysqlUser string `json:"MysqlUser" yaml:"mysql_user"`
	MysqlPwd  string `json:"MysqlPwd" yaml:"mysql_pwd"`
	MysqlHost string `json:"MysqlHost" yaml:"mysql_host"`
	MysqlPort int    `json:"MysqlPort,string" yaml:"mysql_port"`
	MysqlDb   string `json:"MysqlDb" yaml:"mysql_db"`

	OssIsOpen    bool   `json:"OssIsOpen,string" yaml:"oss_isOpen"`
	OssEndpoint  string `json:"OssEndpoint" yaml:"oss_endpoint"`
	OssAccessID  string `json:"OssAccessID" yaml:"oss_accessID"`
	OssAccessKey string `json:"OssAccessKey" yaml:"oss_accessKey"`
	OssDomain    string `json:"OssDomain" yaml:"oss_domain"`
}

func ConfigRead(configFile string) {
	config := new(ConfigMap)
	config.readConfigFromYaml(configFile)
	fmt.Println("Final config: ", config)

	if config.ListenPort == "" {
		panic("init config fail")
	}
	Cfg = config
}

func (_this *ConfigMap) readConfigFromYaml(configFile string) {
	if err := configor.Load(_this, configFile); err != nil {
		fmt.Println("read config yaml err: " + err.Error())
	}
	fmt.Println("Config from yaml file: ", _this)
}
