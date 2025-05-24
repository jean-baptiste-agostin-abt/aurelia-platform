package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() (*gorm.DB, error) {
	user := GetEnv("DB_USER", "root")
	pass := GetEnv("DB_PASS", "password")
	host := GetEnv("DB_HOST", "db")
	name := GetEnv("DB_NAME", "aurelia")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, name)
	return gorm.Open("mysql", dsn)
}
