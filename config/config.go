package config

import (
	"cobra/pkg/db"
	"cobra/pkg/redis"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//LoadManages 加载数据管理服务
func LoadManages() error {

	//启动redis
	if err := redis.Init(); err != nil {
		return err
	}

	//启动db
	if err := db.Init(); err != nil {
		return err
	}

	return nil
}

//MonitorConfig 监听配置
func MonitorConfig() {

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("configuration changes file src: %s", e.Name)
		LoadManages()
	})
}

//Close 关闭资源
func Close() error {

	db.Close()
	redis.Close()
	log.Printf("server colse db redis...")
	return nil
}
