package repository

import (
	"MyGram/model/entity"
	"context"
	"gorm.io/gorm"
)

type SocialMediaRepositoryImpl struct {
	DB *gorm.DB
}

func (smR SocialMediaRepositoryImpl) Create(ctx context.Context, payload entity.SocialMedia) (entity.SocialMedia, error) {
	sm := entity.SocialMedia{
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		UserId:         payload.UserId,
	}
	result := smR.DB.WithContext(ctx).Create(&sm)
	if result.Error != nil {
		return sm, result.Error
	}
	return sm, nil
}

func (smR SocialMediaRepositoryImpl) Update(ctx context.Context, payload entity.SocialMedia) (entity.SocialMedia, error) {
	sm := entity.SocialMedia{
		Id:             payload.Id,
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		UserId:         payload.UserId,
	}
	result := smR.DB.WithContext(ctx).Select("name", "social_media_url").Updates(&sm)

	if result.Error != nil {
		return sm, result.Error
	}
	return sm, nil
}

func (smR SocialMediaRepositoryImpl) Delete(ctx context.Context, payload entity.SocialMedia) (entity.SocialMedia, error) {
	sm := entity.SocialMedia{Id: payload.Id, UserId: payload.UserId}
	result := smR.DB.WithContext(ctx).Delete(&sm)

	if result.Error != nil {
		return sm, result.Error
	}
	return sm, nil
}

func (smR SocialMediaRepositoryImpl) GetById(ctx context.Context, socialMediaId int) (entity.SocialMedia, error) {
	sm := entity.SocialMedia{Id: socialMediaId}
	result := smR.DB.WithContext(ctx).Find(&sm)

	if result.Error != nil {
		return sm, result.Error
	}
	return sm, nil
}

func (smR SocialMediaRepositoryImpl) FindMatch(ctx context.Context, payload entity.SocialMedia) (entity.SocialMedia, error) {
	sm := entity.SocialMedia{Id: payload.Id, UserId: payload.UserId}
	result := smR.DB.WithContext(ctx).Find(&sm)

	if result.Error != nil {
		return sm, result.Error
	}
	return sm, nil
}

func (smR SocialMediaRepositoryImpl) GetByUserId(ctx context.Context, userId int) ([]entity.SocialMedia, error) {
	var sms []entity.SocialMedia
	sm := entity.SocialMedia{UserId: userId}
	result := smR.DB.WithContext(ctx).Where(&sm).Find(&sms)

	if result.Error != nil {
		return sms, result.Error
	}
	return sms, nil
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{DB: db}
}
