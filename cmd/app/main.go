package main

import (
	"botec/internal/dependencies"
	"botec/internal/repositories"
	"botec/internal/service/bot/telegram"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			dependencies.NewConfig,
			dependencies.NewMySql,
			repositories.NewBotRepository,
		),
		fx.Invoke(
			telegram.InitBots,
		),
	).Run()
}
