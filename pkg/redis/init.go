package redis

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var client *redis.Client
var mtx sync.RWMutex

//Init 初始化redis对象
func Init() error {
	mtx.Lock()
	defer mtx.Unlock()

	log.Printf("Initializing Redis Client... %v \n", viper.Get("redis"))
	rds := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password:    viper.GetString("redis.pass"),
		DB:          viper.GetInt("redis.db"),
		ReadTimeout: -1,
	})

	_, err := rds.Ping().Result()
	if err != nil {
		return err
	}

	client = rds
	return nil
}

// Close 关闭redis连接
func Close() {
	if client != nil {
		client.Close()
	}
}

// Client 返回 redis Client
func Client() *redis.Client {
	mtx.RLock()
	defer mtx.RUnlock()
	return client
}
