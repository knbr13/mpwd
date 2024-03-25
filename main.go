package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	r := fiber.New()

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world\n")
	})
	log.Fatal(r.Listen(":3000"))
}
