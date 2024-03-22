package entity

import (
	"time"
)

type User struct {
	Id              int    `gorm:"primaryKey"`
	Username        string `gorm:"uniqueIndex;not null;type:varchar(50)"`
	Email           string `gorm:"uniqueIndex;not null;type:varchar(100)"`
	Password        string `gorm:"not null"`
	Age             int8   `gorm:"not null"`
	ProfileImageUrl string
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
type UserCreateRequest struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	Age             int8   `json:"age" validate:"required,gt=8"`
	ProfileImageUrl string `json:"profile_image_url" validate:"http_url"`
}
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserUpdateRequest struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Age             int8   `json:"age" validate:"required,gt=8"`
	ProfileImageUrl string `json:"profile_image_url" validate:"required,url"`
}

type UserRelationPhoto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserRelationComment struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserRelationSocialMedia struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserCreateResponse struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Age             int8   `json:"age"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type UserReadJwt struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int8   `json:"age"`
}

type UserUpdateResponse struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Age             int8   `json:"age"`
	ProfileImageUrl string `json:"profile_image_url"`
	// Engga pake UpdatedAt
	//UpdateAt        time.Time `json:"update_at"`
}

type TokenLogin struct {
	Token string `json:"token"`
}
