package repository

import (
	"MyGram/model/entity"
	"context"
)

type PhotoRepository interface {
	Create(ctx context.Context, payload entity.Photo) (entity.Photo, error)
	Update(ctx context.Context, payload entity.Photo) (entity.Photo, error)
	GetById(ctx context.Context, photoId int) (entity.Photo, error)
	FindMatch(ctx context.Context, payload entity.Photo) (entity.Photo, error)
	GetAllByUserId(ctx context.Context, userId int) ([]entity.Photo, error)
	Delete(ctx context.Context, payload entity.Photo) error
}
