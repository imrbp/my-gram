package service

import (
	"MyGram/model/entity"
	"context"
)

type SocialMediaService interface {
	Create(context context.Context, payload entity.SocialMediaCreateRequest, auth entity.UserReadJwt) (result entity.SocialMedias, err error)
	Update(context context.Context, payload entity.SocialMediaUpdateRequest, socialMediaId int, auth entity.UserReadJwt) (result entity.SocialMedias, err error)
	Delete(context context.Context, socialMediaId int, auth entity.UserReadJwt) (err error)
	GetAll(context context.Context) (result []entity.SocialMedias, err error)
}
