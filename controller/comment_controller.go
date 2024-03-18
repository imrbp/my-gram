package controller

import (
	"github.com/gofiber/fiber/v2"
)

type CommentController interface {
	Create(ctx *fiber.Ctx) (err error)
	Update(ctx *fiber.Ctx) (err error)
	Delete(ctx *fiber.Ctx) (err error)
	GetAll(ctx *fiber.Ctx) (err error)
}
