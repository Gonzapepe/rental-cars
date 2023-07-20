package repository

import (
	"fmt"
	"math"
	"strings"

	"github.com/Gonzapepe/cars-rental/dtos"
	"github.com/Gonzapepe/cars-rental/model"
	"gorm.io/gorm"
)

// CarRepository the repository of the cars
type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{db: db}
}


func (r *CarRepository) Save(car *model.Car) RepositoryResult {
	err := r.db.Save(car).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: car}
}

func (r *CarRepository) FindAll() RepositoryResult {
	var cars []model.Car

	err := r.db.Find(&cars).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: cars}
}

func (r *CarRepository) FindOneById(id int) RepositoryResult {
	var car model.Car

	err := r.db.Where(&model.Car{Id: id}).Take(&car).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &car}
}

func (r *CarRepository) DeleteOneById(id int) RepositoryResult {
	err := r.db.Delete(&model.Car{Id: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: nil}
}

func (r *CarRepository) Pagination(pagination *dtos.Pagination) (RepositoryResult, int) {
	var cars  []model.Car
	var totalPages int
	var totalRows, fromRow, toRow int64
	totalRows, totalPages, fromRow, toRow = 0, 0, 0, 0

	offset := pagination.Page * pagination.Limit

	// Get data with limit, offset & order
	find := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	// Generate where query
	searchs := pagination.Searchs

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break
			}
		}
	}
	find = find.Find(&cars)

	// Has error finding data
	errFind := find.Error

	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = cars

	// Count all data
	errCount := r.db.Model(&model.Car{}).Count(&totalRows).Error

	if errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}

	pagination.TotalRows = int(totalRows)

	// Calculate total pages 
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from and to row on first page
		fromRow = 1
		toRow = int64(pagination.Limit)
	} else {
		if pagination.Page <= totalPages {
			// calculate from and to row
			fromRow = int64( pagination.Page*pagination.Limit + 1 )
			toRow = int64((pagination.Page + 1) * pagination.Limit)
		}
	}

	if toRow > totalRows {
		toRow = totalRows
	}

	pagination.FromRow = int(fromRow)
	pagination.ToRow = int(toRow)

	return RepositoryResult{Result: pagination}, totalPages
}