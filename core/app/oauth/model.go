package oauth

import (
	"base/core/app/users"
	"time"

	"gorm.io/gorm"
)

type OAuthUser struct {
	users.User     `gorm:"embedded"`
	Provider       string    `gorm:"column:provider"`
	ProviderID     string    `gorm:"column:provider_id"`
	AccessToken    string    `gorm:"column:access_token"`
	OAuthLastLogin time.Time `gorm:"column:oauth_last_login"`
}

func (OAuthUser) TableName() string {
	return "users"
}

type AuthProvider struct {
	gorm.Model
	UserID      uint
	Provider    string
	ProviderID  string
	AccessToken string
	LastLogin   time.Time
}

func (AuthProvider) TableName() string {
	return "auth_providers"
}

// You might want to add OAuth-specific request/response structs here
type OAuthLoginRequest struct {
	Provider    string `json:"provider" binding:"required"`
	AccessToken string `json:"access_token" binding:"required"`
}

type OAuthResponse struct {
	AccessToken string `json:"accessToken"`
	Exp         int64  `json:"exp"`
	Username    string `json:"username"`
	ID          uint   `json:"id"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	LastLogin   string `json:"last_login"`
	Provider    string `json:"provider"`
}
