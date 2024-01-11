package main

import (
	"botec/config"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

const MigrationsLocation = "/usr/src/code/migrations"

var flags = flag.NewFlagSet("goose", flag.ExitOnError)

// Запускает миграции
func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()
	command := args[0]
	arguments := args[1:]

	Config := config.NewConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.Database.Username,
		Config.Database.Password,
		Config.Database.Hostname,
		Config.Database.Name,
	)

	db, err := goose.OpenDBWithDriver("mysql", dsn)

	if err != nil {
		log.Fatalf("Goose: ошибка подключения к БД: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Goose: ошибка закрытия подключения к БД: %v\n", err)
		}
	}()

	if err := goose.Run(command, db, MigrationsLocation, arguments...); err != nil {
		log.Fatalf("Goose %v: %v", command, err)
	}

	fmt.Println(err)
}
