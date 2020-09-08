package main

import (
	"github.com/gofiber/fiber"
	"github.com/qinains/fastergoding" //provides hot reload
)

//An example set up for basic fiber
func main() {
	fastergoding.Run()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Type("html")
		c.Send("<h1> Hello Jon! </h1>")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) {
		c.Type("html")
		if c.Params("name") != "" {
			c.Send("<h1> Hello " + c.Params("name") + "! </h1>")
		} else {
			c.Send("Please Provide A Name to Be Greeted")
		}
	})

	app.Listen(3000)
}
