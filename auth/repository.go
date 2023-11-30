package auth

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) GetItems() ([]string, error) {
	items := []string{}
	return items, r.db.Table("items").Select("name").Scan(&items).Error
}
