package repositories

import (
	"myapi/internal/config"
	"myapi/internal/models"
)

type ItemRepository struct{}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}

func (r *ItemRepository) ListAll() ([]models.Item, error) {
	var itens []models.Item
	if err := config.DB.Find(&itens).Error; err != nil {
		return nil, err
	}
	return itens, nil
}
