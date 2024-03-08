package repository

import (
	"MyGram/model/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, payload entity.User) (entity.User, error)
	Update(ctx context.Context, payload entity.User) (entity.User, error)
	Delete(ctx context.Context, payload entity.User) (err error)
	FindByUsername(ctx context.Context, username string) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
}
