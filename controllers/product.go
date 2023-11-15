package controllers

import "github.com/gofiber/fiber/v2"

func (s Server) CreateProduct(c *fiber.Ctx) error {
	id, err := s.Service.CreateProduct()
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"id": id})
}
