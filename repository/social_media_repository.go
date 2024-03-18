package repository

import (
	"MyGram/model/entity"
	"context"
)

type SocialMediaRepository interface {
	Create(ctx context.Context, payload entity.SocialMedia) (entity.SocialMedia, error)
	Update(ctx context.Context, payload entity.SocialMedia) (entity.SocialMedia, error)
	Delete(ctx context.Context, payload entity.SocialMedia) (entity.SocialMedia, error)
	GetById(ctx context.Context, socialMediaId int) (entity.SocialMedia, error)
	FindMatch(ctx context.Context, payload entity.SocialMedia) (entity.SocialMedia, error)
	GetByUserId(ctx context.Context, userId int) ([]entity.SocialMedia, error)
}
