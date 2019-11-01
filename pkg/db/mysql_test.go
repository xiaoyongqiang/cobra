package db

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
		log.Printf("mysql test err: %v", err)
	}
}

func TestQuery(t *testing.T) {
	if _, err := Engine().Queryx("INSERT INTO bbasic_app_confs SET ac_id = 2, ac_name = ?", "test"); err != nil {
		log.Println(err)
	}
}
