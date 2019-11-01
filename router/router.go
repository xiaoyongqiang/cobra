package router

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

//Run 运行服务
func Run() error {

	// 设置运行模式
	gin.SetMode(viper.GetString("runmode"))

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		log.Println(viper.GetString("runmode"))
	})

	return router.Run(fmt.Sprintf(":%d", viper.GetInt("port")))
}
