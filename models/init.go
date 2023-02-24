package models

import (
	"getcahrzp.cn/define"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db = Init()

var Rdb = InitRedisDB()

func Init() *gorm.DB {
	dsn := define.MysqlDNS + "/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
	return db
}

func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
