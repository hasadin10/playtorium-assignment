package repositories

import (
	"discountmodule/interfaces"

	"gorm.io/gorm"
)

type Repository struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) interfaces.Repository {
	return &Repository{db: db}
}
