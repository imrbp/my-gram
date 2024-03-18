package controller

import (
	"MyGram/model/entity"
	"MyGram/service"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CommentControllerImpl struct {
	Validate       *validator.Validate
	CommentService service.CommentService
}

func NewCommentController(validation *validator.Validate, commentService service.CommentService) CommentController {
	return &CommentControllerImpl{
		Validate:       validation,
		CommentService: commentService,
	}
}

func (cC CommentControllerImpl) Create(ctx *fiber.Ctx) (err error) {
	userReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	commentCreateRequest := entity.CommentCreateRequest{}

	if err := ctx.BodyParser(&commentCreateRequest); err != nil {
		return err
	}
	// TODO: Validate message
	err = cC.Validate.Struct(commentCreateRequest)
	if err != nil {
		return err
	}

	result, err := cC.CommentService.Create(context.Background(), commentCreateRequest, userReadJwt)

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(entity.CommentCreateResponse{
		Id:        result.Id,
		UserId:    result.UserId,
		PhotoId:   result.PhotoId,
		Message:   result.Message,
		CreatedAt: result.CreatedAt,
	})
}

func (cC CommentControllerImpl) Update(ctx *fiber.Ctx) (err error) {
	userReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	commentUpdateRequest := entity.CommentUpdateRequest{}

	if err := ctx.BodyParser(&commentUpdateRequest); err != nil {
		return err
	}
	// TODO: Validation message
	err = cC.Validate.Struct(commentUpdateRequest)
	if err != nil {
		return err
	}

	commentId, err := ctx.ParamsInt(":commentId")
	if err != nil {
		return err
	}
	// TODO:Where is photo id
	photoId, err := ctx.ParamsInt("photoId")
	if err != nil {
		return err
	}
	result, err := cC.CommentService.Update(context.Background(), commentUpdateRequest, commentId, photoId, userReadJwt)

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(entity.CommentUpdateResponse{
		Id:        result.Id,
		UserId:    result.UserId,
		PhotoId:   result.PhotoId,
		Message:   result.Message,
		UpdatedAt: result.UpdatedAt,
	})
}

func (cC CommentControllerImpl) Delete(ctx *fiber.Ctx) (err error) {
	userReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	commentId, err := ctx.ParamsInt("commentId")
	if err != nil {
		return err
	}

	err = cC.CommentService.Delete(context.Background(), commentId, userReadJwt)

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(struct {
		Message string `json:"message"`
	}{Message: "Your comment has been successfully deleted"})
}

func (cC CommentControllerImpl) GetAll(ctx *fiber.Ctx) (err error) {
	userReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	result, err := cC.CommentService.GetAll(context.Background(), userReadJwt)
	if err != nil {
		return err
	}
	var comments []entity.CommentResponse
	for _, comment := range result {
		post := entity.CommentResponse{
			Id:        comment.Id,
			Message:   comment.Message,
			PhotoId:   comment.PhotoId,
			UserId:    comment.UserId,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			User: entity.UserRelationComment{
				Id:       comment.UserId,
				Username: comment.User.Username,
				Email:    comment.User.Email,
			},
			Photo: entity.PhotoRelationComment{
				Id:       comment.Photo.Id,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserId:   comment.Photo.UserId,
			},
		}
		comments = append(comments, post)
	}
	return ctx.Status(fiber.StatusOK).JSON(comments)
}
