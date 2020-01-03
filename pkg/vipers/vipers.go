package vipers

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Vipers 管理集群业务配置
type Vipers struct {
	X *viper.Viper `config:"test"`
	Y *viper.Viper `config:"test2"`
}

var vi = &Vipers{}

// Conf 返回 viper Conf
func Conf() *Vipers {
	return vi
}

//LoadBusinessConfs 加载业务配置
func LoadBusinessConfs(path string) {
	vi.X = viper.New()
	vi.X.AddConfigPath(path)
	vi.X.SetConfigName("test")
	if err := vi.X.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error reading common config file: %v", err))
	}
	vi.X.WatchConfig()
	vi.X.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("configuration changes file src: %s, data: %+v", e.Name, vi.X.AllSettings())
	})

	vi.Y = viper.New()
	vi.Y.AddConfigPath(path)
	vi.Y.SetConfigName("test2")
	if err := vi.Y.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error reading common config file: %v", err))
	}
	vi.Y.WatchConfig()
	vi.Y.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("configuration changes file src: %s, data: %+v", e.Name, vi.Y.AllSettings())
	})
}
