package service

import (
	"MyGram/helper"
	"MyGram/model/entity"
	"MyGram/repository"
	"context"
	"github.com/gofiber/fiber/v2"
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
	result, err := uS.Repository.GetByUsername(ctx, payloadCreate.Username)
	if err != nil {
		return user, fiber.ErrInternalServerError
	}
	if result.Username == payloadCreate.Username {
		return user, fiber.NewError(fiber.StatusNotAcceptable, "Username Already Exists")
	}
	result, err = uS.Repository.GetByEmail(ctx, payloadCreate.Email)
	if err != nil {
		return user, fiber.ErrInternalServerError
	}
	if result.Email == payloadCreate.Email {
		return user, fiber.NewError(fiber.StatusBadRequest, "Email Already Exists")
	}

	hashedPassword, err := helper.HashPassword(payloadCreate.Password)
	if err != nil {
		return user, fiber.NewError(fiber.StatusInternalServerError, "error when hashing Password")
	}
	user = entity.User{
		Username: payloadCreate.Username,
		Email:    payloadCreate.Email,
		Password: hashedPassword,
		Age:      payloadCreate.Age,
	}
	result, err = uS.Repository.Create(ctx, user)
	if err != nil {
		return result, fiber.ErrInternalServerError
	}
	return result, nil
}

func (uS UserServiceImpl) Login(ctx context.Context, payloadLogin entity.UserLoginRequest) (token string, err error) {

	result, err := uS.Repository.GetByEmail(ctx, payloadLogin.Email)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	if result.Email == "" {
		return "", fiber.NewError(fiber.StatusNotFound, "Email doesn't exists in database")
	}

	password, _ := helper.HashPassword(payloadLogin.Password)

	user := entity.User{
		Email:    payloadLogin.Email,
		Password: password,
	}
	result, err = uS.Repository.GetByEmailAndPassword(context.Background(), user)
	if err != nil {
		return "", fiber.NewError(500, "internal server error")
	}
	token, err = helper.GenerateUserToken(result)
	if err != nil {
		return "", fiber.NewError(500, err.Error())
	}
	return token, nil
}

func (uS UserServiceImpl) Update(ctx context.Context, payloadUpdate entity.UserUpdateRequest, auth entity.UserReadJwt) (user entity.User, err error) {
	userRead, err := uS.Repository.FindMatch(ctx, auth)
	if err != nil {
		return user, fiber.NewError(500, "internal server error")
	}

	result, err := uS.Repository.GetByUsername(ctx, payloadUpdate.Username)
	if err != nil {
		return user, fiber.NewError(500, "internal server error")
	}
	if result.Username == payloadUpdate.Username {
		return user, fiber.NewError(404, "Username Already Exists")
	}
	result, err = uS.Repository.GetByEmail(ctx, payloadUpdate.Email)
	if err != nil {
		return user, fiber.NewError(500, "internal server error")
	}
	if result.Email == payloadUpdate.Email {
		return user, fiber.NewError(404, "Email Already Exists")
	}

	userUpdate := entity.User{
		Id:       userRead.Id,
		Email:    payloadUpdate.Email,
		Username: payloadUpdate.Username,
		Age:      userRead.Age,
	}

	result, err = uS.Repository.Update(ctx, userUpdate)
	if err != nil {
		return result, fiber.NewError(500, "error updating")
	}
	return result, nil
}

func (uS UserServiceImpl) Delete(ctx context.Context, auth entity.UserReadJwt) error {
	userRead, err := uS.Repository.FindMatch(ctx, auth)
	if err != nil {
		return fiber.NewError(500, "internal server error")
	}
	userDelete := entity.User{
		Id:       userRead.Id,
		Username: userRead.Username,
		Email:    userRead.Email,
		Age:      userRead.Age,
	}
	err = uS.Repository.Delete(ctx, userDelete)
	if err != nil {
		return fiber.NewError(500, "error updating")
	}
	return nil
}
