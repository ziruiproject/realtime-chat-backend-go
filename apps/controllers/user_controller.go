package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetAll(c *fiber.Ctx) error
}
