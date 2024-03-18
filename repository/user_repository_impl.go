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
	result := uR.DB.WithContext(ctx).Updates(&payload)
	if result.Error != nil {
		return payload, result.Error
	}
	return payload, nil
}

func (uR *UserRepositoryImpl) Delete(ctx context.Context, payload entity.User) (err error) {
	// TODO: Delete With association
	userDeleted := uR.DB.WithContext(ctx).Select(clause.Associations).Delete(&payload)
	if userDeleted.Error != nil {
		return userDeleted.Error
	}
	return nil
}

func (uR *UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	userFind := entity.User{Username: username}
	result := uR.DB.WithContext(ctx).Find(&userFind)
	if result.Error != nil {
		return userFind, result.Error
	}
	return userFind, nil
}

func (uR *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	userFind := entity.User{Email: email}
	result := uR.DB.WithContext(ctx).Find(&userFind)
	if result.Error != nil {
		return userFind, result.Error
	}
	return userFind, nil
}

func (uR *UserRepositoryImpl) FindMatch(ctx context.Context, payload entity.UserReadJwt) (entity.User, error) {
	userFind := entity.User{
		Id:       payload.Id,
		Username: payload.Username,
		Email:    payload.Email,
		Age:      payload.Age,
	}
	result := uR.DB.WithContext(ctx).Find(&userFind)
	if result.Error != nil {
		return userFind, result.Error
	}
	return userFind, nil
}

func (uR *UserRepositoryImpl) GetByEmailAndPassword(ctx context.Context, user entity.User) (entity.User, error) {
	userLogin := entity.User{
		Email:    user.Email,
		Password: user.Password,
	}
	result := uR.DB.WithContext(ctx).Find(&userLogin)
	if result.Error != nil {
		return userLogin, result.Error
	}
	return userLogin, nil
}
