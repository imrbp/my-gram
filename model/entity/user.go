package entity

import (
	"time"
)

type User struct {
	Id        int       `gorm:"primaryKey"`
	Username  string    `gorm:"uniqueIndex;not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Age       int       `gorm:"not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
type UserCreateRequest struct {
	Username string `json:"username" validation:"required"`
	Email    string `json:"email" validation:"required,email"`
	Password string `json:"password" validation:"required,min:6"`
	Age      int    `json:"age" validation:"required,gt:8"`
}
type UserLoginRequest struct {
	Email    string `json:"email" validation:"required,email"`
	Password string `json:"password" validation:"required,min:6"`
}

type UserUpdateRequest struct {
	Username string `json:"username" validation:"required"`
	Email    string `json:"email" validation:"required,email"`
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
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserValidate struct {
	Token string `query:"token"`
}

type UserUpdateResponse struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Age      int       `json:"age"`
	UpdateAt time.Time `json:"update_at"`
}

type TokenLogin struct {
	Token string `json:"token"`
}
