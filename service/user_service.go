package service

import (
	"MyGram/model/entity"
	"context"
)

type UserService interface {
	Register(ctx context.Context, payloadCreate entity.UserCreateRequest) (user entity.User, err error)
	Login(ctx context.Context, payloadCreate entity.UserLoginRequest) (token string, err error)
	Update(ctx context.Context, payloadCreate entity.UserUpdateRequest, auth entity.UserReadJwt) (user entity.User, err error)
	Delete(ctx context.Context, auth entity.UserReadJwt) (err error)
}
