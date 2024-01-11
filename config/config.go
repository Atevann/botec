package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

// NewConfig Чтение конфига
func NewConfig() Config {
	cfg := &Config{}
	_, err := toml.DecodeFile("config/config.toml", cfg)
	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	return *cfg
}

// Config структура конфигурации
type Config struct {
	Database Database
}

// Конфиг базы данных
type Database struct {
	Hostname string
	Name     string
	Username string
	Password string
}
