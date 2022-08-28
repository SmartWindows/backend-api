package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func nada(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func main() {
	app := fiber.New()

	app.Get("/", nada) // Add this

	app.Post("/", nada) // Add this

	app.Put("/update", nada) // Add this

	app.Delete("/delete", nada) // Add this
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
