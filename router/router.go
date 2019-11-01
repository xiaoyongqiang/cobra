package router

import (
	"bytes"
	"cobra/pkg/redis"
	"cobra/tools"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

const REDIS_KEY_Token = "u_token_%s"
const Token_Period = 30 * time.Minute

//Run 运行服务
func Run() error {

	// 设置运行模式
	gin.SetMode(viper.GetString("runmode"))

	router := gin.New()
	router.Use(Authorization)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello go....")
	})

	return router.Run(fmt.Sprintf(":%d", viper.GetInt("port")))
}

//Authorization 中间键验证
func Authorization(c *gin.Context) {
	var sign, uid, token string
	var err error

	// debug或登录接口不验证
	if viper.GetString("runmode") != "debug" {

		// 请求时间必须在30分钟内，防止重放攻击
		stamp, _ := strconv.ParseInt(c.GetHeader("X-TS"), 10, 64)
		ts := time.Unix(stamp, 0)
		if time.Since(ts) > 30*time.Minute || time.Since(ts) < -30*time.Minute {
			log.Printf("close client: request expires, please check both system time %s", c.ClientIP())
			goto END
		}

		uid = c.GetHeader("X-UID")
		sign = c.GetHeader("X-SING")
		if uid == "" || sign == "" {
			goto END
		}

		// 从redis上获取用户token
		if token, err = redis.Client().Get(fmt.Sprintf(REDIS_KEY_Token, uid)).Result(); err != nil {
			log.Printf("get token:%s failed:%v", token, err)
			goto END
		}

		var buf bytes.Buffer
		buf.WriteString(c.GetHeader("X-TS"))
		buf.WriteString(uid)
		buf.WriteString(token)

		if sign != tools.Sign(buf.String()) {
			log.Printf("close client: sign error %v", c.ClientIP())
			goto END
		}
	}

	// Token 延期
	redis.Client().PExpire(fmt.Sprintf(REDIS_KEY_Token, uid), Token_Period)
	c.Next()
	return

END:
	c.Abort()
	c.JSON(http.StatusUnauthorized, gin.H{"status": "401", "msg": "验证失败"})
}
