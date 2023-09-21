// main.go
package main

import (
	"github.com/gofiber/fiber/v2"
)

type Item struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}

var items []Item

func main() {
    app := fiber.New()

    // Create
    app.Post("/items", func(c *fiber.Ctx) error {
        var newItem Item
        if err := c.BodyParser(&newItem); err != nil {
            return err
        }
        newItem.ID = len(items) + 1
        items = append(items, newItem)
        return c.JSON(newItem)
    })

    // Read (Get All)
    app.Get("/items", func(c *fiber.Ctx) error {
        return c.JSON(items)
    })

    // Read (Get One)
    app.Get("/items/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        for _, item := range items {
            if id == string(item.ID) {
                return c.JSON(item)
            }
        }
        return c.Status(404).SendString("Item not found")
    })

    // Update
    app.Put("/items/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        var updatedItem Item
        if err := c.BodyParser(&updatedItem); err != nil {
            return err
        }
        for i, item := range items {
            if id == string(item.ID) {
                items[i] = updatedItem
                return c.JSON(updatedItem)
            }
        }
        return c.Status(404).SendString("Item not found")
    })

    // Delete
    app.Delete("/items/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        for i, item := range items {
            if id == string(item.ID) {
                items = append(items[:i], items[i+1:]...)
                return c.SendString("Item deleted")
            }
        }
        return c.Status(404).SendString("Item not found")
    })

    app.Listen(":3000")
}
