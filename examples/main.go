package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zerolethanh/frlog"
	"log"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})
	app.Put("/login/:uid", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		return c.SendString("User")
	})
	app.Patch("/fly/:from-:to", func(c *fiber.Ctx) error {
		return c.SendString("Fly")
	})

	frlog.PrintAppStacks(app)

	log.Fatalln(app.Listen(":3000"))
}
