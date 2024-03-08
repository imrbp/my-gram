package entity

import (
	"time"
)

type Comment struct {
	Id        int `gorm:"primaryKey"`
	UserId    int
	User      User `gorm:"foreignKey:UserId"`
	PhotoId   int
	Photo     Photo     `gorm:"foreignKey:PhotoId"`
	Message   string    `gorm:"not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type CommentCreateRequest struct {
	PhotoId int    `json:"photo_id"`
	Message string `json:"message" validation:"required"`
}
type CommentCreateResponse struct {
	Id        int       `json:"id"`
	UserId    User      `json:"userId"`
	PhotoId   Photo     `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
type CommentResponse struct {
	Id        int                  `json:"id"`
	Message   string               `json:"message"`
	PhotoId   Photo                `json:"photo_id"`
	UserId    User                 `json:"userId"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	User      UserRelationComment  `json:"User"`
	Photo     PhotoRelationComment `json:"Photo"`
}

type CommentUpdateRequest struct {
	Message string `json:"message" validation:"required"`
}

type CommentUpdateResponse struct {
	Id        int       `json:"id"`
	UserId    User      `json:"userId"`
	PhotoId   Photo     `json:"photo_id"`
	Message   string    `json:"message"`
	UpdatedAt time.Time `json:"updated_at"`
}
