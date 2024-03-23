package service

import (
	"MyGram/model/entity"
	"MyGram/repository"
	"context"
	"github.com/gofiber/fiber/v2"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
}

func NewCommentService(commentRepository repository.CommentRepository) CommentService {
	return &CommentServiceImpl{CommentRepository: commentRepository}
}

func (cS CommentServiceImpl) Create(ctx context.Context, payload entity.CommentCreateRequest, auth entity.UserReadJwt) (result entity.Comment, err error) {

	comment := entity.Comment{
		UserId:  auth.Id,
		PhotoId: payload.PhotoId,
		Message: payload.Message,
	}

	photoFind, err := cS.CommentRepository.GetPhotoById(context.Background(), payload.PhotoId)
	if err != nil {
		return result, err
	}

	// Handling Photo Not found. But when using gorm not acceptable
	if photoFind.Id == 0 {
		return result, fiber.ErrNotFound
	}
	result, err = cS.CommentRepository.Create(ctx, comment)

	if err != nil {
		return result, err
	}
	return result, nil
}

func (cS CommentServiceImpl) Update(ctx context.Context, payload entity.CommentUpdateRequest, commentId int, auth entity.UserReadJwt) (result entity.Comment, err error) {
	comment := entity.Comment{
		Id:     commentId,
		UserId: auth.Id,
	}
	result, err = cS.CommentRepository.GetById(ctx, commentId)
	if err != nil {
		return result, fiber.ErrInternalServerError
	}
	if result.Id != commentId {
		return result, fiber.ErrNotFound
	}
	if result.UserId != auth.Id {
		return result, fiber.ErrUnauthorized
	}
	comment.Message = payload.Message

	result, err = cS.CommentRepository.Update(ctx, comment)
	if err != nil {
		return result, fiber.ErrInternalServerError
	}
	return result, nil

}

func (cS CommentServiceImpl) FindById(ctx context.Context, commentId int, auth entity.UserReadJwt) (result entity.Comment, err error) {
	result, err = cS.CommentRepository.GetById(ctx, commentId)
	if err != nil {
		return result, fiber.ErrInternalServerError
	}
	if result.Id == 0 {
		return result, fiber.ErrNotFound
	}

	return result, nil
}

func (cS CommentServiceImpl) Delete(ctx context.Context, commentId int, auth entity.UserReadJwt) (err error) {

	commentFind, err := cS.CommentRepository.GetById(ctx, commentId)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if commentFind.Id != commentId {
		return fiber.ErrNotFound
	}
	if commentFind.UserId != auth.Id {
		return fiber.ErrUnauthorized
	}

	_, err = cS.CommentRepository.Delete(ctx, commentFind)

	if err != nil {
		return fiber.ErrInternalServerError
	}
	return nil
}

func (cS CommentServiceImpl) GetAll(ctx context.Context, auth entity.UserReadJwt) (comments []entity.Comment, err error) {
	//result, err := cS.CommentRepository.GetByUserId(ctx, auth.Id)

	result, err := cS.CommentRepository.GetAll(ctx)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}
	return result, nil
}
