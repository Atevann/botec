package main

import (
	"botec/internal/dependencies"
	"context"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"log"
	"os"
)

const MigrationsLocation = "/usr/src/code/migrations"

var flags = flag.NewFlagSet("goose", flag.ExitOnError)

// Запускает миграции
func main() {
	fx.New(
		fx.Provide(
			dependencies.NewConfig,
		),
		fx.Invoke(func(config *dependencies.Config, shutdowner fx.Shutdowner) error {
			var db *sql.DB

			defer func() {
				if err := db.Close(); err != nil {
					log.Printf("Goose: ошибка закрытия подключения к БД: %v\n", err)
				}

				log.Println("[Goose] Закрываю подключение к БД")
				shutdowner.Shutdown()
			}()

			flags.Parse(os.Args[1:])
			args := flags.Args()
			command := args[0]
			arguments := args[1:]

			db, err := goose.OpenDBWithDriver("mysql", config.MySql.Dsn)
			if err != nil {
				log.Printf("[Goose] Ошибка подключения к БД: %v\n", err)

				return err
			}

			if err := goose.RunContext(context.Background(), command, db, MigrationsLocation, arguments...); err != nil {
				log.Printf("[Goose] Ошибка при выполнении команды: %v: %v", command, err)

				return err
			}

			return nil
		}),
	).Run()
}
