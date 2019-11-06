package nsq

import (
	"fmt"
	"log"
	"testing"
	"time"

	nsq "github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
)

//Handler 监听消息处理
type Handler struct{}

func init() {
	viper.SetConfigName(".cobra")
	viper.AddConfigPath("../../")
	viper.ReadInConfig()

}

func TestPub(t *testing.T) {
	body := map[string]interface{}{"test": 1, "test2": "ceshi"}
	if err := NsqPubMsg("topic", body); err != nil {
		log.Println("test pub err: ", err)
	}
}

func TestSub(t *testing.T) {
	consumer := NsqConsumer("topic", "channel")
	consumer.AddHandler(&Handler{})
	if err := consumer.ConnectToNSQLookupd(fmt.Sprintf("%s:%d", viper.GetString("nsqlookupd.host"), viper.GetInt("nsqlookupd.port"))); err != nil {
		log.Fatalf("consumer.ConnectToNSQD() fail. Error info: %s\n", err.Error())
	}
	time.Sleep(2 * time.Second)
}

//HandleMessage 处理队列消息
func (h *Handler) HandleMessage(message *nsq.Message) error {
	log.Printf("消费消息: %s", string(message.Body))
	return nil
}
