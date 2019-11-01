package redis

import (
	"log"
	"testing"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".cobra")
	viper.AddConfigPath("../../")
	viper.ReadInConfig()

	if err := Init(); err != nil {
		log.Printf("redis test err: %v", err)
	}
}

func TestGetSet(t *testing.T) {
	Client().Set("test", `{"a":1,"b":3}`, 0)
	log.Printf(Client().Get("test").Val())
}
