package repository

import (
	"MyGram/helper"
	"MyGram/model/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}
func (cR *CommentRepository) Create(ctx context.Context, payload entity.Comment) entity.Comment {
	result := cR.DB.WithContext(ctx).Create(&payload)
	helper.PanicIfError(result.Error)
	return payload
}

func (cR *CommentRepository) Update(ctx context.Context, payload entity.Comment) entity.Comment {
	commentFind := cR.FindById(ctx, payload.Id)
	payloadUpdate := entity.Comment{
		Id:      commentFind.Id,
		Message: payload.Message,
	}
	result := cR.DB.WithContext(ctx).Updates(&payloadUpdate)
	helper.PanicIfError(result.Error)
	return payloadUpdate
}

func (cR *CommentRepository) Delete(ctx context.Context, payload entity.Comment) (err error) {
	// TODO: Delete With association
	commentDeleted := cR.DB.WithContext(ctx).Select(clause.Associations).Delete(payload)
	helper.PanicIfError(commentDeleted.Error)
	return nil
}

func (cR *CommentRepository) FindById(ctx context.Context, commentId int) entity.Comment {
	commentFind := entity.Comment{Id: commentId}
	result := cR.DB.WithContext(ctx).First(&commentFind).Scan(&commentFind)
	helper.PanicIfError(result.Error)
	return commentFind
}
