package repository

import (
	"github.com/Gonzapepe/cars-rental/model"
	"gorm.io/gorm"
)

// MakeRepository the repository of the Make
type MakeRepository struct {
	db *gorm.DB
}

func NewMakeRepository(db *gorm.DB) *MakeRepository {
	return &MakeRepository{db: db}
}

func (r *MakeRepository) Save(make *model.Make) RepositoryResult {
	err := r.db.Save(make).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: make}
}

func (r *MakeRepository) FindAll() RepositoryResult {
	var makers []model.Make

	err := r.db.Find(&makers).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: makers}
}
