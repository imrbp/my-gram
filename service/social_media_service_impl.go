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

func (smS SocialMediaServiceImpl) Create(context context.Context, payload entity.SocialMediaCreateRequest, auth entity.UserReadJwt) (result entity.SocialMedias, err error) {
	socialMedia := entity.SocialMedias{
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		UserId:         auth.Id,
	}
	result, err = smS.SocialMediaRepository.Create(context, socialMedia)
	if err != nil {
		return result, fiber.ErrInternalServerError
	}
	return result, nil
}

func (smS SocialMediaServiceImpl) Update(context context.Context, payload entity.SocialMediaUpdateRequest, socialMediaId int, auth entity.UserReadJwt) (result entity.SocialMedias, err error) {

	socialMedia := entity.SocialMedias{
		Id:             socialMediaId,
		UserId:         auth.Id,
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
	}
	socialMediaFind, err := smS.SocialMediaRepository.GetById(context, socialMediaId)
	if socialMediaFind.Id != socialMediaId {
		return socialMediaFind, fiber.ErrNotFound
	}
	if socialMediaFind.UserId != auth.Id {
		return socialMediaFind, fiber.ErrUnauthorized
	}
	result, err = smS.SocialMediaRepository.Update(context, socialMedia)

	if err != nil {
		return result, fiber.ErrInternalServerError
	}
	return result, nil
}

func (smS SocialMediaServiceImpl) GetAll(context context.Context) (result []entity.SocialMedias, err error) {
	result, err = smS.SocialMediaRepository.GetAll(context)

	if err != nil {
		return result, fiber.ErrInternalServerError
	}
	return result, nil
}
func (smS SocialMediaServiceImpl) Delete(context context.Context, socialMediaId int, auth entity.UserReadJwt) (err error) {
	socialMediaFind, err := smS.SocialMediaRepository.GetById(context, socialMediaId)
	if socialMediaFind.UserId != auth.Id {
		return fiber.ErrUnauthorized
	}

	_, err = smS.SocialMediaRepository.Delete(context, socialMediaFind)

	if err != nil {
		return fiber.ErrInternalServerError
	}
	return nil
}
