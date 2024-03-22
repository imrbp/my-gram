package service

import (
	"MyGram/model/entity"
	"MyGram/repository"
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

type PhotoServiceImpl struct {
	PhotoRepository repository.PhotoRepository
}

func NewPhotoService(photoRepository repository.PhotoRepository) PhotoService {
	return &PhotoServiceImpl{
		PhotoRepository: photoRepository,
	}
}

func (pS PhotoServiceImpl) Post(ctx context.Context, payloadCreate entity.PhotoCreateRequest, auth entity.UserReadJwt) (photo entity.Photo, err error) {
	photoCreate := entity.Photo{
		Title:     payloadCreate.Title,
		Caption:   payloadCreate.Caption,
		PhotoUrl:  payloadCreate.PhotoUrl,
		User:      entity.User{Id: auth.Id},
		UpdatedAt: time.Time{},
		CreatedAt: time.Time{},
	}
	result, err := pS.PhotoRepository.Create(ctx, photoCreate)
	if err != nil {
		return photoCreate, fiber.ErrInternalServerError
	}
	return result, nil
}

func (pS PhotoServiceImpl) Get(ctx context.Context, photoId int, auth entity.UserReadJwt) (photo entity.Photo, err error) {
	photo, err = pS.PhotoRepository.GetById(ctx, photoId)
	if err != nil {
		return photo, err
	}
	if photo.Id == 0 {
		return photo, fiber.ErrNotFound
	}
	if photo.UserId != auth.Id {
		return photo, fiber.ErrUnauthorized
	}
	return photo, nil
}

func (pS PhotoServiceImpl) GetAll(ctx context.Context, auth entity.UserReadJwt) (posts []entity.Photo, err error) {
	posts, err = pS.PhotoRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pS PhotoServiceImpl) Update(ctx context.Context, payloadUpdate entity.PhotoUpdateRequest, photoId int, auth entity.UserReadJwt) (user entity.Photo, err error) {

	result, err := pS.PhotoRepository.GetById(ctx, photoId)

	if err != nil {
		return result, fiber.ErrInternalServerError
	}

	if result.Id == 0 {
		return result, fiber.ErrNotFound
	}
	if result.UserId != auth.Id {
		return result, fiber.ErrUnauthorized
	}

	result, err = pS.PhotoRepository.Update(ctx, entity.Photo{
		Id:       photoId,
		Title:    payloadUpdate.Title,
		Caption:  payloadUpdate.Caption,
		PhotoUrl: payloadUpdate.PhotoUrl,
		UserId:   auth.Id,
	})
	if err != nil {
		return result, fiber.ErrInternalServerError
	}
	return result, nil
}

func (pS PhotoServiceImpl) Delete(ctx context.Context, photoId int, auth entity.UserReadJwt) error {
	photoFind := entity.Photo{
		Id:     photoId,
		UserId: auth.Id,
	}
	photo, err := pS.PhotoRepository.GetById(ctx, photoId)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if photo.Id == 0 {
		return fiber.ErrNotFound
	}

	if photo.UserId != auth.Id {
		return fiber.ErrUnauthorized
	}

	err = pS.PhotoRepository.Delete(ctx, photoFind)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return nil
}
