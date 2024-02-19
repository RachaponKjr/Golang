package main

import (
	"github.com/RachaponKjr/Backend-Golang/controller"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// create Product && update Product && delete Product
	addProduct := app.Group("/api/product")
	addProduct.Post("/create", controller.CreateProduct)
	addProduct.Get("/getall", controller.GetProducts)

	app.Listen(":8080")
}
