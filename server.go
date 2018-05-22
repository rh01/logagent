package main

import (
	"time"

	"github.com/astaxie/beego/logs"
)

func ServerRun() (err error) {
	for true {
		msg := GetOneLine()
		err = senTOKafka(msg)
		if err != nil {
			logs.Error("senfd to kafka failed, err %v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}
func senTOKafka(msg *TextMsg) (err error) {

	// TODO: 完成连接kafka的任务

	logs.Debug("read msg %s, topic: %s", msg.Msg, msg.Topic)
	return
}
