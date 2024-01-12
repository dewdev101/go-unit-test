package services

import "github.com/ThanaponBunchot/go-unit-test/repositories"

type Service interface {
	//method
	CreateProduct() (uint, error)
}

type Dependency struct {
	//Repo
	ProductRepositories repositories.ProductRepositories
}

type service struct {
	productRepo repositories.ProductRepositories
}

func New(d Dependency) Service {
	return service{
		productRepo: d.ProductRepositories,
	}
}

func (s service) CreateProduct() (uint, error) {
	id, err := s.productRepo.CreateProduct()
	if err != nil {
		return 0, err
	}
	return id, nil
}

type promotion struct {
	ID      int
	purchaseMin int
	discount    int
}

func CalculateDiscount(amount int) (promotionPrice int,err error) {

	if amount <= 0 {
		return 0, ErrAmountZero
	}

	if amount >= promotion.purchaseMin {
		discountPrice := amount*(1 - (1*promotion.discount/100))
		return discountPrice,nil
	}

	return promotion{}, nil
}
