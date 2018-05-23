package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
	"github.com/kataras/iris/core/errors"
)

var (
	AppConfig *Config
)

type Config struct {
	LogLevel string
	LogPath  string

	ChanSize    int
	kafkaAddr   string
	collectConf []CollectConf
}

func LoadConf(confType, filename string) (err error) {

	// 按照指定的格式读取配置文件
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Printf("new config failed, err:%v\n", err)
		return
	}

	// 获得配置文件的选项
	AppConfig = &Config{}
	AppConfig.LogLevel = conf.String("logs::log_level")
	if len(AppConfig.LogLevel) == 0 {
		AppConfig.LogLevel = "debug"
	}

	AppConfig.LogPath = conf.String("logs::log_path")
	if len(AppConfig.LogPath) == 0 {
		AppConfig.LogPath = "./logs"
	}

	// 队列大小，或者管道大小
	AppConfig.ChanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		AppConfig.ChanSize = 100
	}

	// kafka地址配置
	AppConfig.kafkaAddr = conf.String("kafka::server_addr")
	if len(AppConfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
		//AppConfig.kafkaAddr = "172.24.1.1"
	}

	err = LoadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed, err %v\n", err)
		return
	}

	return
}

func LoadCollectConf(conf config.Configer) (err error) {

	var cc CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		errors.New("invaild collect::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		errors.New("invaild collect::topic")
		return
	}

	AppConfig.collectConf = append(AppConfig.collectConf, cc)
	return
}
