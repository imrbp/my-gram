package service

import (
	"MyGram/helper"
	"MyGram/model/entity"
	"MyGram/repository"
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

type UserServiceImpl struct {
	Repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &UserServiceImpl{
		Repository: repository,
	}
}
func (uS UserServiceImpl) Register(ctx context.Context, payloadCreate entity.UserCreateRequest) (user entity.User, err error) {
	result, err := uS.Repository.FindByUsername(ctx, payloadCreate.Username)
	if err != nil {
		return user, fiber.NewError(500, "internal server error")
	}
	if result.Username == payloadCreate.Username {
		return user, fiber.NewError(404, "Username Already Exists")
	}
	result, err = uS.Repository.FindByEmail(ctx, payloadCreate.Email)
	if err != nil {
		return user, fiber.NewError(500, "internal server error")
	}
	if result.Email == payloadCreate.Email {
		return user, fiber.NewError(404, "Email Already Exists")
	}

	hashedPassword, err := helper.HashPassword(payloadCreate.Password)
	if err != nil {
		return user, fiber.NewError(500, "error when hashing Password")
	}
	user = entity.User{
		Username: payloadCreate.Username,
		Email:    payloadCreate.Email,
		Password: hashedPassword,
		Age:      payloadCreate.Age,
	}
	result, err = uS.Repository.Create(ctx, user)
	if err != nil {
		return result, fiber.NewError(404, "error when creating User")
	}
	return result, nil
}

func (uS UserServiceImpl) Login(ctx context.Context, payloadLogin entity.UserLoginRequest) (token string) {
	hash
	user := entity.User{
		Email:     payloadLogin.Email,
		Email:     "",
		Password:  "",
		Age:       0,
		UpdatedAt: time.Time{},
		CreatedAt: time.Time{},
	}
}
