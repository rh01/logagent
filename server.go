package main

import (
	"time"
	"./kafka"
	"github.com/astaxie/beego/logs"
)

func ServerRun() (err error) {
	for {
		msg := GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("senfd to kafka failed, err %v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *TextMsg) (err error) {
	err = kafka.SendTOKafka(msg.Msg, msg.Topic)

	return
}
