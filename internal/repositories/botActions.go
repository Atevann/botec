package repositories

import (
	"botec/internal/dependencies"
	"botec/internal/models"
	"fmt"
)

// BotActionsRepository Репозиторий экшенов бота
type BotActionsRepository struct {
	db *dependencies.MySql
}

// NewBotActionsRepository Инициализирует репозиторий экщенов
func NewBotActionsRepository(database *dependencies.MySql) *BotActionsRepository {
	return &BotActionsRepository{db: database}
}

// GetNextAction Получить следующий экшен
func (repo *BotActionsRepository) GetNextAction(
	botId uint,
	parentActionId uint,
	parentCondition string,
) (*models.BotAction, error) {
	var firstAction models.BotAction
	statement := fmt.Sprintf("bot_id = %v AND parent_id = %v", botId, parentActionId)

	if len(parentCondition) != 0 {
		statement = fmt.Sprintf("%v AND parent_condition = %v", parentCondition)
	}

	result := repo.db.First(&firstAction, statement)

	if result.Error != nil {
		return nil, result.Error
	}

	return &firstAction, nil
}
