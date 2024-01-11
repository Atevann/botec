package dependencies

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// MySql Структура подключения к БД
type MySql struct {
	*gorm.DB
}

// NewMySql Подключение к MySql через gorm
func NewMySql(config *Config) (*MySql, error) {
	db, err := gorm.Open(
		mysql.Open(config.MySql.Dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

	if err != nil {
		log.Printf("[Gorm] Ошибка подключения к БД: %s", err)
	}

	return &MySql{DB: db}, err
}
