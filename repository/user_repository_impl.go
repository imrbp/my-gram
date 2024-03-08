package repository

import (
	"MyGram/model/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (uR *UserRepositoryImpl) Create(ctx context.Context, payload entity.User) (entity.User, error) {
	result := uR.DB.WithContext(ctx).Create(&payload)
	if result.Error != nil {
		return payload, result.Error
	}
	return payload, nil
}

func (uR *UserRepositoryImpl) Update(ctx context.Context, payload entity.User) (entity.User, error) {
	userFind, err := uR.FindByUsername(ctx, payload.Username)
	if err != nil {
		return userFind, err
	}
	payloadUpdate := entity.User{
		Id:       userFind.Id,
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Age:      payload.Age,
	}
	result := uR.DB.WithContext(ctx).Updates(&payloadUpdate)
	if result.Error != nil {
		return payloadUpdate, result.Error
	}
	return payloadUpdate, nil
}

func (uR *UserRepositoryImpl) Delete(ctx context.Context, payload entity.User) (err error) {
	// TODO: Delete With association
	userDeleted := uR.DB.WithContext(ctx).Select(clause.Associations).Delete(payload)
	if userDeleted.Error != nil {
		return userDeleted.Error
	}
	return nil
}

func (uR *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (entity.User, error) {
	userFind := entity.User{Username: username}
	result := uR.DB.WithContext(ctx).First(&userFind).Scan(&userFind)
	if result.Error != nil {
		return userFind, result.Error
	}
	return userFind, nil
}

func (uR *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	userFind := entity.User{Email: email}
	result := uR.DB.WithContext(ctx).First(&userFind).Scan(&userFind)
	if result.Error != nil {
		return userFind, result.Error
	}
	return userFind, nil
}

//func (uR *UserRepositoryImpl) FindMatch(ctx context.Context, entities entity.User) entity.User {
//	findMatch := entity.User{}
//	result := uR.DB.WithContext(ctx).First(&entities).Scan(&findMatch)
//	helper.PanicIfError(result.Error)
//	return findMatch
//}
