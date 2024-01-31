package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID int `json:"id"`
	Tile string `json:"title"`
	Done bool `json:"done"`
	Body string `json:"body"`
}

func main() {
	fmt.Println("Hello World")

	// Assing a new Fiber instance to app variable
	// := returns a pointer to a new instance of Fiber
	app := fiber.New()

	todos := []Todo{}

	// The error is saying we can also return an error
	app.Get("/healthcheck", func(fiberContext *fiber.Ctx) error {
		return fiberContext.SendString("Hello, World!")
	})

	app.Post("/api/todos", func(fiberContext *fiber.Ctx) error {
		todo := &Todo{}

		if err := fiberContext.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return fiberContext.JSON(todo)
	})

	app.Patch("/api/todos/:id/done", func(fiberContext *fiber.Ctx) error {
		id, err := fiberContext.ParamsInt("id")

		if err != nil {
			return fiberContext.Status(401).SendString("Invalid ID")
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Done = true
				break;
			}
		}

		return fiberContext.JSON(todos)
	})

	// If app.Listen() throws an error, log it
	log.Fatal(app.Listen(":4000"))
}