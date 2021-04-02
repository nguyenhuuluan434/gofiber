package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{Id: 1, Title: "Walk", Completed: false},
	{Id: 2, Title: "Eat", Completed: false},
}

func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		}})
}

func CreateTodo(c *fiber.Ctx) error {
	type Request struct {
		Title string `json:"title"`
	}

	var body Request
	err := c.BodyParser(&body)

	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,
			"message": "Cannot parse JSON"})
	}
	todo := &Todo{
		Id:        len(todos) + 1,
		Title:     body.Title,
		Completed: false,
	}
	todos = append(todos, todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}
