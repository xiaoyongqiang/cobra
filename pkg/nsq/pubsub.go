package nsq

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	nsq "github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
)

//NsqConsumer Nsq消费者
func NsqConsumer(topic, channel string) *nsq.Consumer {

	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	consumer, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		log.Fatalf("nsq.NewConsumer() fail. Error info: %s\n", err.Error())
	}

	consumer.SetLogger(log.New(os.Stderr, "", log.Flags()), nsq.LogLevelWarning)
	return consumer
}

//NsqProducer 生产者对象
func NsqProducer() (*nsq.Producer, error) {
	cfg := nsq.NewConfig()
	return nsq.NewProducer(fmt.Sprintf("%s:%d", viper.GetString("nsq.host"), viper.GetInt("nsq.port")), cfg)
}

//NsqPubMsg 推送消息队列
func NsqPubMsg(topic string, body map[string]interface{}) error {
	w, err := NsqProducer()
	defer w.Stop()
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return w.Publish(topic, bytes)
}
