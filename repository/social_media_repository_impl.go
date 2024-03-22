package repository

import (
	"MyGram/model/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepositoryImpl struct {
	DB *gorm.DB
}

func (smR SocialMediaRepositoryImpl) Create(ctx context.Context, payload entity.SocialMedias) (entity.SocialMedias, error) {
	sm := entity.SocialMedias{
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

func (smR SocialMediaRepositoryImpl) Update(ctx context.Context, payload entity.SocialMedias) (entity.SocialMedias, error) {
	sm := entity.SocialMedias{
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

func (smR SocialMediaRepositoryImpl) Delete(ctx context.Context, payload entity.SocialMedias) (entity.SocialMedias, error) {
	sm := entity.SocialMedias{Id: payload.Id, UserId: payload.UserId}
	result := smR.DB.WithContext(ctx).Delete(&sm)

	if result.Error != nil {
		return sm, result.Error
	}
	return sm, nil
}

func (smR SocialMediaRepositoryImpl) GetById(ctx context.Context, socialMediaId int) (entity.SocialMedias, error) {
	sm := entity.SocialMedias{}
	result := smR.DB.WithContext(ctx).Where(entity.SocialMedias{Id: socialMediaId}).Find(&sm)

	if result.Error != nil {
		return sm, result.Error
	}
	return sm, nil
}

func (smR SocialMediaRepositoryImpl) FindMatch(ctx context.Context, payload entity.SocialMedias) (entity.SocialMedias, error) {
	sm := entity.SocialMedias{}
	result := smR.DB.WithContext(ctx).Where(entity.SocialMedias{Id: payload.Id, UserId: payload.Id}).Find(&sm)

	if result.Error != nil {
		return sm, result.Error
	}
	return sm, nil
}

func (smR SocialMediaRepositoryImpl) GetByUserId(ctx context.Context, userId int) ([]entity.SocialMedias, error) {
	var sms []entity.SocialMedias
	result := smR.DB.WithContext(ctx).Where(entity.SocialMedias{UserId: userId}).Find(&sms)

	if result.Error != nil {
		return sms, result.Error
	}
	return sms, nil
}

func (smR SocialMediaRepositoryImpl) GetAll(ctx context.Context) ([]entity.SocialMedias, error) {
	var sms []entity.SocialMedias
	result := smR.DB.WithContext(ctx).Preload(clause.Associations).Model(entity.SocialMedias{}).Find(&sms)

	if result.Error != nil {
		return sms, result.Error
	}
	return sms, nil
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{DB: db}
}
