package controller

import (
	"MyGram/helper"
	"MyGram/model/entity"
	"MyGram/service"
	"context"
	"github.com/gofiber/fiber/v2"
)

type PhotoControllerImpl struct {
	Service    service.PhotoService
	XValidator *helper.Validator
}

func NewPhotoController(photoService service.PhotoService, validator *helper.Validator) PhotoController {
	return &PhotoControllerImpl{
		Service:    photoService,
		XValidator: validator,
	}
}

func (pC PhotoControllerImpl) Post(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)
	ctx.Accepts("application/json")

	photoCreateRequest := entity.PhotoCreateRequest{}

	err = pC.XValidator.ParseBody(ctx, &photoCreateRequest)
	if err != nil {
		return err
	}

	post, err := pC.Service.Post(context.Background(), photoCreateRequest, UserReadJwt)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(entity.PhotoCreateResponse{
		Id:       post.Id,
		Title:    post.Title,
		Caption:  post.Caption,
		PhotoUrl: post.PhotoUrl,
		UserId:   post.UserId,
		//CreatedAt: post.CreatedAt,
	})
}

func (pC PhotoControllerImpl) Get(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	photoId, err := ctx.ParamsInt("photoId")
	if err != nil {
		return err
	}
	result, err := pC.Service.Get(context.Background(), photoId, UserReadJwt)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(entity.PhotoResponse{
		Id:       result.Id,
		Title:    result.Title,
		Caption:  result.Caption,
		PhotoUrl: result.PhotoUrl,
		UserId:   result.UserId,
		User: entity.UserRelationPhoto{
			Username: UserReadJwt.Username,
			Email:    UserReadJwt.Email,
		},
	})

}

func (pC PhotoControllerImpl) GetAll(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	result, err := pC.Service.GetAll(context.Background(), UserReadJwt)
	var posts []entity.PhotoResponse
	//TODO: Not Efficient
	if len(result) == 0 {
		return ctx.Status(fiber.StatusOK).JSON([]string{})
	}
	for _, photo := range result {
		post := entity.PhotoResponse{
			Id:       photo.Id,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoUrl: photo.PhotoUrl,
			UserId:   photo.UserId,
			//UpdatedAt: photo.UpdatedAt,
			//CreatedAt: photo.CreatedAt,
			User: entity.UserRelationPhoto{
				Username: photo.User.Username,
				Email:    photo.User.Email,
			},
		}
		posts = append(posts, post)
	}
	return ctx.Status(fiber.StatusOK).JSON(posts)

}

func (pC PhotoControllerImpl) Update(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	photoUpdateRequest := entity.PhotoUpdateRequest{}

	err = pC.XValidator.ParseBody(ctx, &photoUpdateRequest)
	if err != nil {
		return err
	}

	photoId, err := ctx.ParamsInt("photoId")
	if err != nil {
		return err
	}
	result, err := pC.Service.Update(context.Background(), photoUpdateRequest, photoId, UserReadJwt)

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(entity.PhotoUpdateResponse{
		Id:       result.Id,
		Title:    result.Title,
		Caption:  result.Caption,
		PhotoUrl: result.PhotoUrl,
		UserId:   result.UserId,
		//UpdatedAt: result.UpdatedAt,
	})
}

func (pC PhotoControllerImpl) Delete(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	photoId, err := ctx.ParamsInt("photoId")
	if err != nil {
		return err
	}
	err = pC.Service.Delete(context.Background(), photoId, UserReadJwt)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(struct {
		Message string `json:"message"`
	}{"Your Photo has been successfully deleted"})
}
