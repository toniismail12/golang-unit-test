// main.go
package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	err := app.Listen(":5001")
	if err != nil {
		fmt.Println(err)
	}
}
