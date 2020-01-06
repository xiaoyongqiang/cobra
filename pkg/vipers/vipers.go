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

//Viperface 对像封装
// type Viperface struct {
// 	V       *viper.Viper
// 	Tags    string
// 	DirPath string
// }

var vi = &Vipers{}

// Conf 返回 viper Conf
func Conf() *Vipers {
	return vi
}

//LoadBusinessConfs 加载业务配置
func LoadBusinessConfs(path string) {
	vi.X = viper.New()
	LoadingObjects(path, "test", vi.X)

	vi.Y = viper.New()
	LoadingObjects(path, "test2", vi.Y)
}

//LoadingObjects 加载到独立对象
func LoadingObjects(path string /*文件夹路径*/, file string /*文件名字*/, v *viper.Viper /*独立配置对象*/) {

	v.AddConfigPath(path)
	v.SetConfigName(file)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error reading common config file: %v", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("configuration changes file src: %s, data: %+v", e.Name, v.AllSettings())
	})
}

// //GetViperTag 获取对应tag
// func GetViperTag(a *viper.Viper) {
// 	s := reflect.TypeOf(*a) //通过反射获取type定义
// 	for i := 0; i < s.NumField(); i++ {
// 		fmt.Println(s.Field(i).Tag.Get("config")) //将tag输出出来
// 	}
// }
