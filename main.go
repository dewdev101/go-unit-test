package main

import (
	"github.com/ThanaponBunchot/go-unit-test/config"
	"github.com/ThanaponBunchot/go-unit-test/controllers"
	"github.com/ThanaponBunchot/go-unit-test/repositories"
	"github.com/ThanaponBunchot/go-unit-test/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := config.ConnectDatabase()
	// dbTest := config.ConnectDatabaseTest()
	productRepository := repositories.NewProductRepository(repositories.ProductDatbase{GormDB: db})

	sv := services.New(services.Dependency{
		ProductRepositories: productRepository,
	})

	s := controllers.New(controllers.Dependency{
		App:     app,
		Service: sv,
	})
	s.Route()

	// s.App.Get("v1/", func(c *fiber.Ctx) error {
	// 	productRepository.CreateProduct()
	// 	return nil
	// })

	app.Listen(":9000")
}
