package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"tudo/util"
)

const (
	CacheNil = redis.Nil
)

type DBConfig struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	Dbname string `json:"db_name"`
	Param  string `json:"param"`
}

type CacheConfig struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	DB     int    `json:"db"`
	Prefix string `json:"prefix"`
}

var DBConfigObj DBConfig
var CacheConfigObj CacheConfig

var DB *gorm.DB
var Cache *redis.Client

func DBInit(path string) {
	if err := util.ReadJSON(path, &DBConfigObj); err != nil {
		panic(err)
	}

	config := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?%v",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		DBConfigObj.Host,
		DBConfigObj.Port,
		DBConfigObj.Dbname,
		DBConfigObj.Param,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&User{}, &UserInfo{}, &Activity{})
	if err != nil {
		panic(err)
	}

	return
}

func CacheInit(path string) {
	if err := util.ReadJSON(path, &CacheConfigObj); err != nil {
		panic(err)
	}

	Cache = redis.NewClient(&redis.Options{
		Addr:     CacheConfigObj.Host + ":" + CacheConfigObj.Port,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       CacheConfigObj.DB,
	})

	if _, err := Cache.Ping().Result(); err != nil {
		panic(err)
	}
	return
}
