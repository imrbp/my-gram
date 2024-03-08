package entity

import (
	"time"
)

type Photo struct {
	Id        int    `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Caption   string
	PhotoUrl  string `gorm:"not null"`
	UserId    int
	User      User      `gorm:"foreignKey:UserId"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type PhotoCreateRequest struct {
	Title    string `json:"title" validation:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validation:"required"`
}

type PhotoResponse struct {
	Id        int               `json:"id"`
	Title     string            `json:"title"`
	Caption   string            `json:"caption"`
	PhotoUrl  string            `json:"photo_url"`
	UserId    User              `json:"user_id"`
	UpdatedAt time.Time         `json:"updated_at"`
	CreatedAt time.Time         `json:"created_at"`
	User      UserRelationPhoto `json:"User"`
}

type PhotoUpdateRequest struct {
	Id        int       `json:"id"  validation:"required"`
	Title     string    `json:"title"  validation:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"  validation:"required"`
	UserId    User      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoRelationComment struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   User   `json:"user_id"`
}
