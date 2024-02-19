package controller

import (
	"time"

	"github.com/RachaponKjr/Backend-Golang/model"
	"github.com/gofiber/fiber/v2"
)

var products = []model.Product{
	{ID: 1, Name: "Product 1", Quantity: 10, Price: 100, CreateTime: time.Now(), UpdateTime: time.Now()},
	{ID: 2, Name: "Product 2", Quantity: 20, Price: 200, CreateTime: time.Now(), UpdateTime: time.Now()}}

func CreateProduct(c *fiber.Ctx) error {

	product := new(model.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	product.CreateTime = time.Now()
	product.UpdateTime = time.Now()
	products = append(products, *product)
	return c.JSON(product)
}

func GetProducts(c *fiber.Ctx) error {
	return c.JSON(products)
}
