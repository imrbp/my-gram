package repository

import (
	"MyGram/model/entity"
	"context"
)

type SocialMediaRepository interface {
	Create(ctx context.Context, payload entity.SocialMedias) (entity.SocialMedias, error)
	Update(ctx context.Context, payload entity.SocialMedias) (entity.SocialMedias, error)
	Delete(ctx context.Context, payload entity.SocialMedias) (entity.SocialMedias, error)
	GetById(ctx context.Context, socialMediaId int) (entity.SocialMedias, error)
	GetAll(ctx context.Context) ([]entity.SocialMedias, error)
	FindMatch(ctx context.Context, payload entity.SocialMedias) (entity.SocialMedias, error)
	GetByUserId(ctx context.Context, userId int) ([]entity.SocialMedias, error)
}
