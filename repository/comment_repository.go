package repository

import (
	"MyGram/model/entity"
	"context"
)

type CommentRepository interface {
	Create(context context.Context, comment entity.Comment) (entity.Comment, error)
	GetById(context context.Context, commentId int) (entity.Comment, error)
	FindMatch(context context.Context, comment entity.Comment) (entity.Comment, error)
	Update(context context.Context, comment entity.Comment) (entity.Comment, error)
	Delete(context context.Context, comment entity.Comment) (entity.Comment, error)
	GetByUserId(context context.Context, userId int) ([]entity.Comment, error)
}
