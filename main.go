package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kwe/go-rest-api/hwb"
)

func main() {
	config := fiber.Config{
		ServerHeader: "kwe",
		ETag:         true,
	}
	app := fiber.New(config)

	api := app.Group("/api")

	hwb.Register(api)

	log.Fatal(app.Listen(":3000"))
}
