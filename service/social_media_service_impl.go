package service

import (
	"MyGram/model/entity"
	"MyGram/repository"
	"context"
	"github.com/gofiber/fiber/v2"
)

type SocialMediaServiceImpl struct {
	SocialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository repository.SocialMediaRepository) SocialMediaService {
	return &SocialMediaServiceImpl{SocialMediaRepository: socialMediaRepository}
}

func (smS SocialMediaServiceImpl) Create(context context.Context, payload entity.SocialMediaCreateRequest, auth entity.UserReadJwt) (result entity.SocialMedia, err error) {
	socialMedia := entity.SocialMedia{
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		UserId:         auth.Id,
	}
	result, err = smS.SocialMediaRepository.Create(context, socialMedia)
	if err != nil {
		return socialMedia, fiber.ErrInternalServerError
	}
	return socialMedia, nil
}

func (smS SocialMediaServiceImpl) Update(context context.Context, payload entity.SocialMediaUpdateRequest, socialMediaId int, auth entity.UserReadJwt) (result entity.SocialMedia, err error) {

	socialMediaFind, err := smS.SocialMediaRepository.GetById(context, socialMediaId)
	if socialMediaFind.UserId != auth.Id {
		return socialMediaFind, fiber.ErrForbidden
	}
	socialMediaFind.SocialMediaUrl = payload.SocialMediaUrl
	socialMediaFind.Name = payload.Name
	result, err = smS.SocialMediaRepository.Update(context, socialMediaFind)

	if err != nil {
		return socialMediaFind, fiber.ErrInternalServerError
	}
	return result, nil
}

func (smS SocialMediaServiceImpl) GetAll(context context.Context, auth entity.UserReadJwt) (result []entity.SocialMedia, err error) {
	result, err = smS.SocialMediaRepository.GetByUserId(context, auth.Id)

	if err != nil {
		return result, fiber.ErrInternalServerError
	}
	return result, nil
}
func (smS SocialMediaServiceImpl) Delete(context context.Context, socialMediaId int, auth entity.UserReadJwt) (err error) {
	socialMediaFind, err := smS.SocialMediaRepository.GetById(context, socialMediaId)
	if socialMediaFind.UserId != auth.Id {
		return fiber.ErrForbidden
	}

	_, err = smS.SocialMediaRepository.Delete(context, socialMediaFind)

	if err != nil {
		return fiber.ErrInternalServerError
	}
	return nil
}
