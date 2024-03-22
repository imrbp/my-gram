package controller

import "C"
import (
	"MyGram/helper"
	"MyGram/model/entity"
	"MyGram/service"
	"context"
	"github.com/gofiber/fiber/v2"
)

type SocialMediaControllerImpl struct {
	XValidation        *helper.Validator
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaController(validation *helper.Validator, socialMediaService service.SocialMediaService) SocialMediaController {
	return &SocialMediaControllerImpl{
		XValidation:        validation,
		SocialMediaService: socialMediaService,
	}
}

func (smC SocialMediaControllerImpl) Create(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	socialMediaCreateRequest := entity.SocialMediaCreateRequest{}

	err = smC.XValidation.ParseBody(ctx, &socialMediaCreateRequest)
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
		//CreatedAt:      result.CreatedAt,
	})
}

func (smC SocialMediaControllerImpl) Update(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	socialMediaUpdateRequest := entity.SocialMediaUpdateRequest{}
	err = smC.XValidation.ParseBody(ctx, &socialMediaUpdateRequest)
	if err != nil {
		return err
	}

	socialMediaId, err := ctx.ParamsInt("socialMediaId")
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
		//UpdatedAt:      result.UpdatedAt,
	})
}

func (smC SocialMediaControllerImpl) Delete(ctx *fiber.Ctx) (err error) {
	UserReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	socialMediaId, err := ctx.ParamsInt("socialMediaId")
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
	if UserReadJwt.Id == 0 {
		return fiber.ErrUnauthorized
	}
	result, err := smC.SocialMediaService.GetAll(context.Background())
	if err != nil {
		return err
	}

	var socialMedias []entity.ItemSocialMedia
	if len(result) == 0 {
		return ctx.Status(fiber.StatusOK).JSON([]string{})
	}
	for _, media := range result {
		socialMedia := entity.ItemSocialMedia{
			Id:             media.Id,
			Name:           media.Name,
			SocialMediaUrl: media.SocialMediaUrl,
			UserId:         media.UserId,
			//UpdatedAt:      media.UpdatedAt,
			//CreatedAt:      media.CreatedAt,
			User: entity.UserRelationSocialMedia{
				Username: media.User.Username,
				Email:    media.User.Email,
			},
		}
		socialMedias = append(socialMedias, socialMedia)
	}

	return ctx.Status(fiber.StatusOK).JSON(socialMedias)
}
