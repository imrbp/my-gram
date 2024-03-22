package controller

import (
	"github.com/gofiber/fiber/v2"
)

type PhotoController interface {
	Post(ctx *fiber.Ctx) (err error)
	Get(ctx *fiber.Ctx) (err error)
	GetAll(ctx *fiber.Ctx) (err error)
	Update(ctx *fiber.Ctx) (err error)
	Delete(ctx *fiber.Ctx) (err error)
}
