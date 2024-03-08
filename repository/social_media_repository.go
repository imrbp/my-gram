package repository

import (
	"MyGram/helper"
	"MyGram/model/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepository struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{DB: db}
}
func (cR *SocialMediaRepository) Create(ctx context.Context, payload entity.SocialMedia) entity.SocialMedia {
	result := cR.DB.WithContext(ctx).Create(&payload)
	helper.PanicIfError(result.Error)
	return payload
}

func (cR *SocialMediaRepository) Update(ctx context.Context, payload entity.SocialMedia) entity.SocialMedia {
	socialMediaFind := cR.FindById(ctx, payload.Id)
	payloadUpdate := entity.SocialMedia{
		Id:             socialMediaFind.Id,
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
	}
	result := cR.DB.WithContext(ctx).Updates(&payloadUpdate)
	helper.PanicIfError(result.Error)
	return payloadUpdate
}

func (cR *SocialMediaRepository) Delete(ctx context.Context, payload entity.SocialMedia) (err error) {
	// TODO: Delete With association
	socialMediaDeleted := cR.DB.WithContext(ctx).Select(clause.Associations).Delete(payload)
	helper.PanicIfError(socialMediaDeleted.Error)
	return nil
}

func (cR *SocialMediaRepository) FindById(ctx context.Context, commentId int) entity.SocialMedia {
	commentFind := entity.SocialMedia{Id: commentId}
	result := cR.DB.WithContext(ctx).First(&commentFind).Scan(&commentFind)
	helper.PanicIfError(result.Error)
	return commentFind
}
