# Goose - инструмент для миграций https://github.com/pressly/goose
# Создание миграции:   make goose GOOSE_CMD="create migration_name sql"
# Применение миграции: make goose GOOSE_CMD="up"
# Откат миграции:      make goose GOOSE_CMD="down"
GOOSE_CMD = status
goose:
	docker run --rm \
            --name goose \
            --network container:mysql \
            -v $(shell pwd)/:/usr/src/code \
            -w /usr/src/code \
            golang \
            go run cmd/cli/goose/main.go $(GOOSE_CMD)