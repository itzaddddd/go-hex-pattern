package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itzaddddd/go-hex/adapter"
	"github.com/itzaddddd/go-hex/core"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	orderRepo := adapter.NewOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapter.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	db.AutoMigrate(&core.Order{})

	app.Listen(":8080")
}
