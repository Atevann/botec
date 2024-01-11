package dependencies

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

// Config структура конфигурации
type Config struct {
	MySql MySqlConfig
}

// MySql конфиг БД
type MySqlConfig struct {
	Hostname string
	Name     string
	Username string
	Password string
	Dsn      string
}

// NewConfig сборка конфига
func NewConfig() (*Config, error) {
	cfg := &Config{}

	_, err := toml.DecodeFile("config/config.toml", cfg)

	cfg.MySql.generateDsn()

	if err != nil {
		log.Printf("Ошибка чтения конфига: %s", err)
	}

	return cfg, err
}

// generateDsn Генерирует DSN
func (config *MySqlConfig) generateDsn() {
	config.Dsn = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Hostname,
		config.Name,
	)
}
