package controller

import (
	"MyGram/model/entity"
	"MyGram/service"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	Service  service.UserService
	Validate *validator.Validate
}

func NewUserController(userService service.UserService, validator *validator.Validate) UserController {
	return &UserControllerImpl{
		Service:  userService,
		Validate: validator,
	}
}

func (uC UserControllerImpl) Register(ctx *fiber.Ctx) (err error) {
	ctx.Accepts("application/json")
	userRegisterRequest := entity.UserCreateRequest{}

	if err := ctx.BodyParser(&userRegisterRequest); err != nil {
		return err
	}
	// TODO: Validation message
	err = uC.Validate.Struct(userRegisterRequest)
	if err != nil {
		return err
	}

	result, err := uC.Service.Register(context.Background(), userRegisterRequest)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(entity.UserCreateResponse{
		Id:       result.Id,
		Username: result.Username,
		Email:    result.Email,
		Age:      result.Age,
	})
}

func (uC UserControllerImpl) Login(ctx *fiber.Ctx) (err error) {

	ctx.Accepts("application/json")
	userLogin := entity.UserLoginRequest{}

	if err := ctx.BodyParser(&userLogin); err != nil {
		return err
	}
	// TODO: Validation message
	err = uC.Validate.Struct(userLogin)
	if err != nil {
		return err
	}

	token := uC.Service.Login(context.Background(), userLogin)
	return ctx.Status(fiber.StatusOK).JSON(entity.TokenLogin{
		Token: token,
	};
}
