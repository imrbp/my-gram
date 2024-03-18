package repository

import (
	"MyGram/model/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, payload entity.User) (entity.User, error)
	Update(ctx context.Context, payload entity.User) (entity.User, error)
	Delete(ctx context.Context, payload entity.User) (err error)
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	FindMatch(ctx context.Context, payload entity.UserReadJwt) (entity.User, error)
	GetByEmailAndPassword(ctx context.Context, payload entity.User) (entity.User, error)
}
