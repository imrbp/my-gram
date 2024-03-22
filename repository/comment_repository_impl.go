package repository

import (
	"MyGram/model/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (cR CommentRepositoryImpl) Create(context context.Context, comment entity.Comment) (entity.Comment, error) {
	payload := entity.Comment{
		UserId:  comment.UserId,
		PhotoId: comment.PhotoId,
		Message: comment.Message,
	}
	result := cR.DB.WithContext(context).Create(&payload)

	if result.Error != nil {
		return payload, result.Error
	}
	return payload, nil
}

func (cR CommentRepositoryImpl) GetById(context context.Context, commentId int) (entity.Comment, error) {
	payload := entity.Comment{}
	result := cR.DB.WithContext(context).Where(entity.Comment{Id: commentId}).Find(&payload)
	if result.Error != nil {
		return payload, result.Error
	}
	return payload, nil
}

func (cR CommentRepositoryImpl) GetPhotoById(context context.Context, photoId int) (entity.Photo, error) {
	payload := entity.Photo{}
	result := cR.DB.WithContext(context).Where(entity.Comment{Id: photoId}).Find(&payload)
	if result.Error != nil {
		return payload, result.Error
	}
	return payload, nil
}

func (cR CommentRepositoryImpl) GetAll(context context.Context) ([]entity.Comment, error) {
	var photos []entity.Comment
	result := cR.DB.WithContext(context).Preload(clause.Associations).Model(entity.Comment{}).Find(&photos)
	if result.Error != nil {
		return photos, result.Error
	}
	return photos, nil

}

func (cR CommentRepositoryImpl) Update(context context.Context, comment entity.Comment) (entity.Comment, error) {
	payload := entity.Comment{
		Id:      comment.Id,
		UserId:  comment.UserId,
		PhotoId: comment.PhotoId,
		Message: comment.Message,
	}
	result := cR.DB.WithContext(context).Select("message").Updates(&payload)
	if result.Error != nil {
		return payload, result.Error
	}
	return payload, nil
}

func (cR CommentRepositoryImpl) Delete(context context.Context, comment entity.Comment) (entity.Comment, error) {
	payload := entity.Comment{Id: comment.Id, UserId: comment.UserId}
	result := cR.DB.WithContext(context).Delete(&payload)
	if result.Error != nil {
		return payload, result.Error
	}
	return payload, nil
}

func (cR CommentRepositoryImpl) GetByUserId(context context.Context, userId int) ([]entity.Comment, error) {
	var comments []entity.Comment

	result := cR.DB.WithContext(context).Where(entity.Comment{UserId: userId}).Scan(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil

}
