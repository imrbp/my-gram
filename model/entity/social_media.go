package entity

import (
	"time"
)

type SocialMedia struct {
	Id             int    `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	SocialMediaUrl string `gorm:"not null"`
	UserId         int
	User           User      `gorm:"foreignKey:UserId"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

type SocialMediaCreateRequest struct {
	Name           string `json:"name" validation:"required"`
	SocialMediaUrl string `json:"social_media_url" validation:"required,url"`
}

type SocialMediaCreateResponse struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}
type GetSocialMedia struct {
	SocialMedias []ItemSocialMedia `json:"social_medias"`
}

type ItemSocialMedia struct {
	Id             int                     `json:"id"`
	Name           string                  `json:"name"`
	SocialMediaUrl string                  `json:"social_media_url"`
	UserId         int                     `json:"user_id"`
	UpdatedAt      time.Time               `json:"updated_at"`
	CreatedAt      time.Time               `json:"created_at"`
	User           UserRelationSocialMedia `json:"User"`
}
type SocialMediaUpdateRequest struct {
	Name           string `json:"name" validation:"required"`
	SocialMediaUrl string `json:"social_media_url" validation:"required"`
}

type SocialMediaUpdatedResponse struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}
