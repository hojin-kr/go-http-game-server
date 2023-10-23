package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	_ "github.com/hojin-kr/go-http-game-server/cmd/docs"
	AccountModel "github.com/hojin-kr/go-http-game-server/cmd/model/account"
)

// @title go-http-game-server API
// @version 1.0.0
// @description go-http-game-server API
// @host localhost:3000
// @BasePath /
func main() {
	log.Println(uuid.New())
	app := fiber.New()

	// Match any route
	app.Use(func(c *fiber.Ctx) error {
		log.Println("🥇 First handler")
		return c.Next()
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")

	// APIs for go-http-server

	// Account APIs

	// GET GetAccountByUUID
	v1.Get("/account/:uuid", func(c *fiber.Ctx) error {
		uuid := c.Params("uuid")
		account := AccountModel.GetByUUID(uuid)
		return c.JSON(account)
	})
	// todo: POST Update Account APPLE PLATFORM Token by UUID
	log.Fatal(app.Listen(":3000"))
}
