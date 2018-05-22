package main

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {

	// 读取配置
	filename := "./conf/logcollect.conf"
	err := LoadConf("ini", filename)
	if err != nil {
		fmt.Println("load conf failed, err: %v", err)
		panic(err)
		return
	}

	//初始化配置
	err = InitLogger()
	if err != nil {
		fmt.Println("load conf failed, err: %v", err)
		panic(err)
		return
	}

	logs.Debug("initial success!")
	logs.Debug("logs filename: %v", AppConfig)

	// 初始化tailf配置
	// Refer this document： https://github.com/hpcloud/tail
	err = InitTail(AppConfig.collectConf, AppConfig.ChanSize)
	if err != nil {
		logs.Error("init tail failed, error %v", err)
		//fmt.Println("init tail failed, err %v", err)
		return
	}

	logs.Debug("init all success...")

	// 真正的业务逻辑处理函数
	err = ServerRun()
	if err != nil {
		logs.Error("server run failed ...")
		return
	}

	logs.Info("program exited...")

}
