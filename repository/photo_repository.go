package repository

import (
	"MyGram/helper"
	"MyGram/model/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{DB: db}
}
func (pR *PhotoRepository) Create(ctx context.Context, payload entity.Photo) entity.Photo {
	result := pR.DB.WithContext(ctx).Create(&payload)
	helper.PanicIfError(result.Error)
	return payload
}

func (pR *PhotoRepository) Update(ctx context.Context, payload entity.Photo) entity.Photo {
	photoFind := pR.FindById(ctx, payload.Id)
	payloadUpdate := entity.Photo{
		Id:       photoFind.Id,
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoUrl,
	}
	result := pR.DB.WithContext(ctx).Updates(&payloadUpdate)
	helper.PanicIfError(result.Error)
	return payloadUpdate
}

func (pR *PhotoRepository) Delete(ctx context.Context, payload entity.Photo) (err error) {
	// TODO: Delete With association
	photoDeleted := pR.DB.WithContext(ctx).Select(clause.Associations).Delete(payload)
	helper.PanicIfError(photoDeleted.Error)
	return nil
}

func (pR *PhotoRepository) FindById(ctx context.Context, photoId int) entity.Photo {
	photoFind := entity.Photo{Id: photoId}
	result := pR.DB.WithContext(ctx).First(&photoFind).Scan(&photoFind)
	helper.PanicIfError(result.Error)
	return photoFind
}
