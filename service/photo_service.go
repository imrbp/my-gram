package service

import (
	"MyGram/model/entity"
	"context"
)

type PhotoService interface {
	Post(ctx context.Context, payloadCreate entity.PhotoCreateRequest, auth entity.UserReadJwt) (photo entity.Photo, err error)
	Get(ctx context.Context, photoId int, auth entity.UserReadJwt) (photo entity.Photo, err error)
	GetAll(ctx context.Context, auth entity.UserReadJwt) (posts []entity.Photo, err error)
	Update(ctx context.Context, payloadCreate entity.PhotoUpdateRequest, photoId int, auth entity.UserReadJwt) (user entity.Photo, err error)
	Delete(ctx context.Context, photoId int, auth entity.UserReadJwt) (err error)
}
