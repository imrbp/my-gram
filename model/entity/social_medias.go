package entity

import (
	"time"
)

type SocialMedias struct {
	Id             int       `gorm:"primaryKey"`
	Name           string    `gorm:"not null;type:varchar(50)"`
	SocialMediaUrl string    `gorm:"not null"`
	UserId         int       `gorm:"not null"`
	User           User      `gorm:"foreignKey:UserId"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

type SocialMediaCreateRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required,url"`
}

type SocialMediaCreateResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         int    `json:"user_id"`

	// Engga pake semuanya yang di bawah ini
	//CreatedAt      time.Time `json:"created_at"`
}

// Engga pake semuanya yang di bawah ini
//type GetSocialMedia struct {
//	SocialMedias []ItemSocialMedia `json:"social_medias"`
//}

type ItemSocialMedia struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         int    `json:"user_id"`

	// Engga pake semuanya yang di bawah ini
	//UpdatedAt      time.Time               `json:"updated_at"`
	//CreatedAt      time.Time               `json:"created_at"`
	User UserRelationSocialMedia `json:"User"`
}
type SocialMediaUpdateRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required,url"`
}

type SocialMediaUpdatedResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         int    `json:"user_id"`

	// Engga pake semuanya yang di bawah ini
	//UpdatedAt      time.Time `json:"updated_at"`
}
