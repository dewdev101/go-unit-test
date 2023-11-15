package repositories

import (
	"github.com/ThanaponBunchot/go-unit-test/model"
	"gorm.io/gorm"
)

type ProductRepositoriesTest interface {
	//method
	CreateProduct() (uint, error)
}

type productRepositoryTest struct {
	// repo
	GormDB *gorm.DB
}

type ProductDatbaseTest struct {
	GormDB *gorm.DB
}

func NewProductRepositoryTest(d ProductDatbaseTest) ProductRepositoriesTest {
	return productRepositoryTest{
		GormDB: d.GormDB,
	}
}

func (repo productRepositoryTest) CreateProduct() (uint, error) {
	input := model.Product{
		Name:  "ItemTetst",
		Price: 100.45,
	}
	res := repo.GormDB.Create(&input)
	if res.Error != nil {
		return 0, res.Error
	}
	return 0, nil
}
