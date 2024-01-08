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

type Config struct {
	Bot Bot
}

type Bot struct {
	Token string
}
