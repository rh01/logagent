package main

import (
	"time"

	"github.com/apex/log"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}

type TextMsg struct {
	Msg   string
	Topic string
}

type TailObjMgr struct {
	tailsObjs []*TailObj
	msgChan   chan *TextMsg
}

var (
	tailObjMgr *TailObjMgr
)

func GetOneLine() (msg *TextMsg) {

	msg = <-tailObjMgr.msgChan
	return
}

func InitTail(conf []CollectConf, chanSize int) (err error) {

	// 初始化管道
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	// 容错处理
	if len(conf) == 0 {
		logs.Error("invaild config for log collect, conf:%v", conf)
		return
	}

	for _, v := range conf {

		obj := &TailObj{
			conf: v,
		}

		tails, tailError := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,
			Follow:    true,
			Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:      true,
		})

		if tailError != nil {

			log.Errorf("tailf occurs errors, error: %v", tailError)
			return tailError
		}

		obj.tail = tails
		tailObjMgr.tailsObjs = append(tailObjMgr.tailsObjs, obj)

		go ReadFromTail(obj)

	}

	return
}

func ReadFromTail(tailObj *TailObj) {

	for true {
		line, ok := <-tailObj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen, filename:%s\n", tailObj.tail.Filename)
			time.Sleep(1000 * time.Millisecond)
			continue
		}

		textMsg := &TextMsg{
			Msg:   line.Text,
			Topic: tailObj.conf.Topic,
		}

		tailObjMgr.msgChan <- textMsg

	}
}
