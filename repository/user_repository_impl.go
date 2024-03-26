package repository

import (
	"MyGram/model/entity"
	"context"
	"gorm.io/gorm"
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

	_ = uR.DB.WithContext(ctx).Where(entity.Photo{UserId: payload.Id}).Delete(&entity.Photo{})
	_ = uR.DB.WithContext(ctx).Where(entity.Comment{UserId: payload.Id}).Delete(&entity.Comment{})
	_ = uR.DB.WithContext(ctx).Where(entity.SocialMedias{UserId: payload.Id}).Delete(&entity.SocialMedias{})

	userDeleted := uR.DB.WithContext(ctx).Where(entity.User{Id: payload.Id}).Delete(&entity.User{Id: payload.Id})
	if userDeleted.Error != nil {
		return userDeleted.Error
	}
	return nil
}

func (uR *UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	userFind := entity.User{}
	result := uR.DB.WithContext(ctx).Where(&entity.User{Username: username}).Find(&userFind)
	if result.Error != nil {
		return userFind, result.Error
	}
	return userFind, nil
}

func (uR *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	userFind := entity.User{}
	result := uR.DB.WithContext(ctx).Where(&entity.User{Email: email}).Find(&userFind)
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
	result := uR.DB.WithContext(ctx).Where(&userFind)
	if result.Error != nil {
		return userFind, result.Error
	}
	return userFind, nil
}

func (uR *UserRepositoryImpl) GetByEmailAndPassword(ctx context.Context, user entity.User) (entity.User, error) {
	userLogin := entity.User{}
	result := uR.DB.WithContext(ctx).Where(&entity.User{Email: user.Email, Password: user.Password}).First(&userLogin)
	if result.Error != nil {
		return userLogin, result.Error
	}
	return userLogin, nil
}
