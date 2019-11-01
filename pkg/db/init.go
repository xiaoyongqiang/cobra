package db

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var db *sqlx.DB
var mtx sync.RWMutex

//Init 数据库初始化
func Init() error {
	mtx.Lock()
	defer mtx.Unlock()

	log.Printf("Initializing Mysql... %v \n", viper.Get("db"))
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("db.user"),
		viper.GetString("db.pass"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.name"),
		viper.GetString("db.charset"),
	)

	mysql, err := sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}

	mysql.SetMaxIdleConns(viper.GetInt("db.idleConns"))
	mysql.SetMaxOpenConns(viper.GetInt("db.openConns"))

	db = mysql
	return nil
}

// Close 关闭数据库连接
func Close() {
	if db != nil {
		db.Close()
	}
}

// Engine mysql对象
func Engine() *sqlx.DB {
	mtx.RLock()
	defer mtx.RUnlock()
	return db
}
