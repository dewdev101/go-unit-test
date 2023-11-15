package repositories

import (
	"github.com/ThanaponBunchot/go-unit-test/model"
	"gorm.io/gorm"
)

type ProductRepositories interface {
	//method
	CreateProduct() (uint, error)
}

type productRepository struct {
	// repo
	GormDB *gorm.DB
}

type ProductDatbase struct {
	GormDB *gorm.DB
}

func NewProductRepository(d ProductDatbase) ProductRepositories {
	return productRepository{
		GormDB: d.GormDB,
	}
}

func (repo productRepository) CreateProduct() (uint, error) {
	input := model.Product{
		Name:  "Item1",
		Price: 100.45,
	}
	res := repo.GormDB.Create(&input)
	if res.Error != nil {
		return 0, res.Error
	}
	return 0, nil
}
