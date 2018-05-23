package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) (err error){
	// 初始化配置
	config := sarama.NewConfig()
	// Ack
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 分区负载均衡
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 初始化生产者
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {

		logs.Error("init kafka producer failed, err: " ,err)
		return
	}

	logs.Debug("init kafka success ...")
	return
}

func SendTOKafka(data, topic string) (err error) {

	msg := &sarama.ProducerMessage{}
	// 消息内容
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)


	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, err:%v, data:%v, topic:%v\n", err,data,topic)
		return
	}

	logs.Debug("send success, pid %v, offset:%v, topic %v\n", pid, offset, topic)
	return

}