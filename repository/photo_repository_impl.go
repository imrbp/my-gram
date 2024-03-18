package repository

import (
	"MyGram/model/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepositoryImpl struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{DB: db}
}

func (pR PhotoRepositoryImpl) Create(ctx context.Context, payload entity.Photo) (entity.Photo, error) {
	result := pR.DB.WithContext(ctx).Create(&payload)
	if result.Error != nil {
		return payload, result.Error
	}
	return payload, nil
}

func (pR PhotoRepositoryImpl) Update(ctx context.Context, payload entity.Photo) (entity.Photo, error) {
	result := pR.DB.WithContext(ctx).Updates(&payload)
	if result.Error != nil {
		return payload, result.Error
	}
	return payload, nil
}

func (pR PhotoRepositoryImpl) GetById(ctx context.Context, photoId int) (entity.Photo, error) {
	photoFind := entity.Photo{Id: photoId}
	result := pR.DB.WithContext(ctx).Find(&photoFind)
	if result.Error != nil {
		return photoFind, result.Error
	}
	return photoFind, nil
}

func (pR PhotoRepositoryImpl) FindMatch(ctx context.Context, payload entity.Photo) (entity.Photo, error) {
	result := pR.DB.WithContext(ctx).Find(&payload)
	if result.Error != nil {
		return payload, result.Error
	}
	return payload, nil
}

func (pR PhotoRepositoryImpl) GetAllByUserId(ctx context.Context, userId int) ([]entity.Photo, error) {
	var photos []entity.Photo

	result := pR.DB.WithContext(ctx).Where("user_id", userId).Find(&photos)
	if result.Error != nil {
		return nil, result.Error
	}
	return photos, nil
}

func (pR PhotoRepositoryImpl) Delete(ctx context.Context, payload entity.Photo) error {
	photoDeleted := pR.DB.WithContext(ctx).Select(clause.Associations).Delete(&payload)
	if photoDeleted.Error != nil {
		return photoDeleted.Error
	}
	return nil

}
