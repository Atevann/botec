package repositories

import (
	"botec/internal/dependencies"
	"botec/internal/models"
)

// BotRepository Репозиторий ботов
type BotRepository struct {
	db *dependencies.MySql
}

// NewBotRepository Инициализация репозитория ботов
func NewBotRepository(database *dependencies.MySql) *BotRepository {
	return &BotRepository{db: database}
}

// GetOneById Получение одного бота по ID
func (repo *BotRepository) GetOneById(id uint) (*models.Bot, error) {
	var bot models.Bot

	result := repo.db.First(&bot, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &bot, nil
}

// GetAll Получение всех ботов
func (repo *BotRepository) GetAll() ([]*models.Bot, error) {
	var bots []*models.Bot

	result := repo.db.Find(&bots)

	if result.Error != nil {
		return nil, result.Error
	}

	return bots, nil
}

// Update Обновление записи бота
func (repo *BotRepository) Update(bot *models.Bot) error {
	result := repo.db.Save(bot)

	return result.Error
}

// Delete Удаление записи бота
func (repo *BotRepository) Delete(bot *models.Bot) error {
	result := repo.db.Delete(bot)

	return result.Error
}
