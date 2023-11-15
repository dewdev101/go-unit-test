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
