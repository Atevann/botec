package main

import (
	"botec/config"
	"botec/internal/service/bot/telegram"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	Config := config.NewConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.Database.Username,
		Config.Database.Password,
		Config.Database.Hostname,
		Config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных %v", err)
	}

	err = telegram.InitBots(db)
	if err != nil {
		log.Println(err)
	}
}
