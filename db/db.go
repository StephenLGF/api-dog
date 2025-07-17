// db/db.go
package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
}

func GetDB() *gorm.DB {
	once.Do(func() {
		configFile, err := os.ReadFile("config.json")
		if err != nil {
			log.Fatal(err)
		}
		// 解析配置文件
		var config Config
		err = json.Unmarshal(configFile, &config)
		if err != nil {
			log.Fatal(err)
		}

		// 连接数据库
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.Name)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
		}
	})
	return db
}
