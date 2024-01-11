package repositories

import (
	"botec/internal/models"
	"gorm.io/gorm"
)

// BotRepository Структура репозитория
type BotRepository struct {
	db *gorm.DB
}

// NewBotRepository Инициализация репозитория ботов
func NewBotRepository(database *gorm.DB) *BotRepository {
	return &BotRepository{db: database}
}

// GetOneById Получение одного бота по ID
func (br *BotRepository) GetOneById(id uint) (*models.Bot, error) {
	var bot models.Bot

	result := br.db.First(&bot, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &bot, nil
}

// GetAll Получение всех ботов
func (br *BotRepository) GetAll() ([]*models.Bot, error) {
	var bots []*models.Bot

	result := br.db.Find(&bots)

	if result.Error != nil {
		return nil, result.Error
	}

	return bots, nil
}

// Update Обновление записи бота
func (br *BotRepository) Update(bot *models.Bot) error {
	result := br.db.Save(bot)

	return result.Error
}

// Delete Удаление записи бота
func (br *BotRepository) Delete(bot *models.Bot) error {
	result := br.db.Delete(bot)

	return result.Error
}
