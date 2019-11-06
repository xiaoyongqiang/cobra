package testsub

import (
	"fmt"

	nsq "github.com/nsqio/go-nsq"
)

//TestHandler 监听消息处理
type TestHandler struct{}

//HandleMessage 处理队列消息
func (h *TestHandler) HandleMessage(message *nsq.Message) error {

	fmt.Printf("%s", string(message.Body))
	return nil
}
