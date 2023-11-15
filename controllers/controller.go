package controllers

import (
	"github.com/ThanaponBunchot/go-unit-test/services"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App     *fiber.App
	Service services.Service
}
type Dependency struct {
	App     *fiber.App
	Service services.Service
}

func New(d Dependency) Server {
	return Server{
		App:     d.App,
		Service: d.Service,
	}
}
