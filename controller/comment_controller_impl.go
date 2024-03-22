package controller

import (
	"MyGram/helper"
	"MyGram/model/entity"
	"MyGram/service"
	"context"
	"github.com/gofiber/fiber/v2"
)

type CommentControllerImpl struct {
	XValidator     *helper.Validator
	CommentService service.CommentService
}

func NewCommentController(validation *helper.Validator, commentService service.CommentService) CommentController {
	return &CommentControllerImpl{
		XValidator:     validation,
		CommentService: commentService,
	}
}

func (cC CommentControllerImpl) Create(ctx *fiber.Ctx) (err error) {
	userReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	commentCreateRequest := entity.CommentCreateRequest{}

	err = cC.XValidator.ParseBody(ctx, &commentCreateRequest)
	if err != nil {
		return err
	}

	result, err := cC.CommentService.Create(context.Background(), commentCreateRequest, userReadJwt)

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(entity.CommentCreateResponse{
		Id:      result.Id,
		UserId:  result.UserId,
		PhotoId: result.PhotoId,
		Message: result.Message,
		//CreatedAt: result.CreatedAt,
	})
}

func (cC CommentControllerImpl) Update(ctx *fiber.Ctx) (err error) {
	userReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	commentUpdateRequest := entity.CommentUpdateRequest{}

	err = cC.XValidator.ParseBody(ctx, &commentUpdateRequest)
	if err != nil {
		return err
	}

	commentId, err := ctx.ParamsInt("commentId")
	if err != nil {
		return err
	}
	result, err := cC.CommentService.Update(context.Background(), commentUpdateRequest, commentId, userReadJwt)

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(entity.CommentUpdateResponse{
		Id:      result.Id,
		UserId:  result.UserId,
		PhotoId: result.PhotoId,
		Message: result.Message,
		//UpdatedAt: result.UpdatedAt,
	})
}

func (cC CommentControllerImpl) Get(ctx *fiber.Ctx) (err error) {
	userReadJwt := ctx.Locals("userRead").(entity.UserReadJwt)

	commentId, err := ctx.ParamsInt("commentId")
	if err != nil {
		return err
	}

	result, err := cC.CommentService.FindById(context.Background(), commentId, userReadJwt)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(entity.CommentResponse{
		Id:      result.Id,
		Message: result.Message,
		PhotoId: result.PhotoId,
		UserId:  result.UserId,
		User: entity.UserRelationComment{
			Id:       result.User.Id,
			Username: result.User.Username,
			Email:    result.User.Email,
		},
		Photo: entity.PhotoRelationComment{
			Id:       result.Photo.Id,
			Title:    result.Photo.Title,
			Caption:  result.Photo.Caption,
			PhotoUrl: result.Photo.PhotoUrl,
			UserId:   result.Photo.UserId,
		},
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

	//TODO: No efficient
	if len(result) == 0 {
		return ctx.Status(fiber.StatusOK).JSON([]string{})
	}
	for _, comment := range result {
		post := entity.CommentResponse{
			Id:      comment.Id,
			Message: comment.Message,
			PhotoId: comment.PhotoId,
			UserId:  comment.UserId,
			//CreatedAt: comment.CreatedAt,
			//UpdatedAt: comment.UpdatedAt,
			User: entity.UserRelationComment{
				//Userid expose
				Id:       comment.UserId,
				Username: comment.User.Username,
				Email:    comment.User.Email,
			},
			Photo: entity.PhotoRelationComment{
				//photo Id Expose
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
