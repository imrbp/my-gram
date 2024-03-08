package service

import (
	"MyGram/model/entity"
	"context"
)

type UserService interface {
	Register(ctx context.Context, payloadCreate entity.UserCreateRequest) (user entity.User, err error)
	Login(ctx context.Context, payloadCreate entity.UserLoginRequest) (token string)
}
