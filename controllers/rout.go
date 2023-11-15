package controllers

func (s Server) Route() {
	v1 := s.App.Group("v2")
	v1.Get("/", s.CreateProduct)
}
