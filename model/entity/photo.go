package entity

import (
	"time"
)

type Photo struct {
	Id        int       `gorm:"primaryKey"`
	Title     string    `gorm:"not null;type:varchar(100)"`
	Caption   string    `gorm:"type:varchar(200)"`
	PhotoUrl  string    `gorm:"not null"`
	UserId    int       `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserId"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type PhotoCreateRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required,url"`
}

type PhotoCreateResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
	// Engga pake Created At
	//CreatedAt time.Time `json:"created_at"`
}

type PhotoResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
	// Engga pake Created At and UpdatedAt
	//UpdatedAt time.Time         `json:"updated_at"`
	//CreatedAt time.Time         `json:"created_at"`

	// dan engga pake User
	User UserRelationPhoto `json:"User"`
}

type PhotoUpdateRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required,url"`
}

type PhotoUpdateResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`

	// Engga pake UpdatedAt
	//UpdatedAt time.Time `json:"updated_at"`
}

type PhotoRelationComment struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}
