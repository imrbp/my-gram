package entity

import (
	"time"
)

type Comment struct {
	Id        int       `gorm:"primaryKey"`
	UserId    int       `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserId"`
	PhotoId   int       `gorm:"not null"`
	Photo     Photo     `gorm:"foreignKey:PhotoId"`
	Message   string    `gorm:"not null;type:varchar(200)"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type CommentCreateRequest struct {
	PhotoId int    `json:"photo_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}
type CommentCreateResponse struct {
	Id      int    `json:"id"`
	UserId  int    `json:"userId"`
	PhotoId int    `json:"photo_id"`
	Message string `json:"message"`

	// Engga pake semuanya yang di bawah ini
	//CreatedAt time.Time `json:"created_at"`
}
type CommentResponse struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	PhotoId int    `json:"photo_id"`
	UserId  int    `json:"userId"`
	// Engga pake semuanya yang di bawah ini
	//CreatedAt time.Time            `json:"created_at"`
	//UpdatedAt time.Time            `json:"updated_at"`
	User  UserRelationComment  `json:"User"`
	Photo PhotoRelationComment `json:"Photo"`
}

type CommentUpdateRequest struct {
	Message string `json:"message" validate:"required"`
}

type CommentUpdateResponse struct {
	Id      int    `json:"id"`
	UserId  int    `json:"userId"`
	PhotoId int    `json:"photo_id"`
	Message string `json:"message"`

	// Engga pake semuanya yang di bawah ini
	//UpdatedAt time.Time `json:"updated_at"`
}
