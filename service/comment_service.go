package service

import (
	"MyGram/model/entity"
	"context"
)

type CommentService interface {
	Create(ctx context.Context, payload entity.CommentCreateRequest, auth entity.UserReadJwt) (result entity.Comment, err error)
	Update(ctx context.Context, payload entity.CommentUpdateRequest, commentId int, auth entity.UserReadJwt) (result entity.Comment, err error)
	FindById(ctx context.Context, commentId int, auth entity.UserReadJwt) (result entity.Comment, err error)
	Delete(ctx context.Context, commentId int, auth entity.UserReadJwt) (err error)
	GetAll(ctx context.Context, auth entity.UserReadJwt) (comments []entity.Comment, err error)
}
