package hwb

import "github.com/gofiber/fiber/v2"

func GetIndex(c *fiber.Ctx) error {
	return c.SendString("Hwb homepage")
}

func search(c *fiber.Ctx) error {
	return c.SendString("Hwb search")
}

func login(c *fiber.Ctx) error {
	return c.SendString("Hwb login")
}

func Register(router fiber.Router) {
	hwbRouter := router.Group("/hwb")
	hwbRouter.Get("/", GetIndex)
	hwbRouter.Get("/search", search)
	hwbRouter.Get("/login", login)
}
