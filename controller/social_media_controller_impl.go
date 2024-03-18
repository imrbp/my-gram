package controller

import "C"
import (
	"MyGram/model/entity"
	"MyGram/service"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SocialMediaControllerImpl struct {
	Validate           *validator.Validate
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaController(validation *validator.Validate, socialMediaService service.SocialMediaService) SocialMediaController {
	return &SocialMediaControllerImpl{
		Validate:           validation,
		SocialMediaService: socialMediaService,
	}
}

func (smC SocialMediaControllerImpl) Create(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	socialMediaCreateRequest := entity.SocialMediaCreateRequest{}

	if err := ctx.BodyParser(&socialMediaCreateRequest); err != nil {
		return err
	}
	// TODO: Validation message
	err = smC.Validate.Struct(socialMediaCreateRequest)
	if err != nil {
		return err
	}

	result, err := smC.SocialMediaService.Create(context.Background(), socialMediaCreateRequest, UserReadJwt)

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(entity.SocialMediaCreateResponse{
		Id:             result.Id,
		Name:           result.Name,
		SocialMediaUrl: result.SocialMediaUrl,
		UserId:         result.UserId,
		CreatedAt:      result.CreatedAt,
	})
}

func (smC SocialMediaControllerImpl) Update(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	socialMediaUpdateRequest := entity.SocialMediaUpdateRequest{}

	if err := ctx.BodyParser(&socialMediaUpdateRequest); err != nil {
		return err
	}
	// TODO: Validation message
	err = smC.Validate.Struct(socialMediaUpdateRequest)
	if err != nil {
		return err
	}

	socialMediaId, err := ctx.ParamsInt(":socialMediaId")
	if err != nil {
		return err
	}
	result, err := smC.SocialMediaService.Update(context.Background(), socialMediaUpdateRequest, socialMediaId, UserReadJwt)

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(entity.SocialMediaUpdatedResponse{
		Id:             result.Id,
		Name:           result.Name,
		SocialMediaUrl: result.SocialMediaUrl,
		UserId:         result.UserId,
		UpdatedAt:      result.UpdatedAt,
	})
}

func (smC SocialMediaControllerImpl) Delete(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	socialMediaId, err := ctx.ParamsInt(":socialMediaId")
	if err != nil {
		return err
	}
	err = smC.SocialMediaService.Delete(context.Background(), socialMediaId, UserReadJwt)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(struct {
		Message string `json:"message"`
	}{"Your Social Media has been successfully deleted"})
}

func (smC SocialMediaControllerImpl) GetAll(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	result, err := smC.SocialMediaService.GetAll(context.Background(), UserReadJwt)
	if err != nil {
		return err
	}

	var socialMedias []entity.ItemSocialMedia
	for _, media := range result {
		socialMedia := entity.ItemSocialMedia{
			Id:             media.Id,
			Name:           media.Name,
			SocialMediaUrl: media.SocialMediaUrl,
			UserId:         media.UserId,
			UpdatedAt:      media.UpdatedAt,
			CreatedAt:      media.CreatedAt,
			User: entity.UserRelationSocialMedia{
				Username: media.User.Username,
				Email:    media.User.Email,
			},
		}
		socialMedias = append(socialMedias, socialMedia)
	}

	return ctx.Status(fiber.StatusOK).JSON(entity.GetSocialMedia{SocialMedias: socialMedias})
}
