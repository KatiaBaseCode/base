package auth

import (
	"base/core/app/users"
	"time"
)

type AuthUser struct {
	users.User       `gorm:"embedded"`
	LastLogin        *time.Time `gorm:"column:last_login"`
	ResetToken       string     `gorm:"column:reset_token"`
	ResetTokenExpiry *time.Time `gorm:"column:reset_token_expiry"`
}

func (AuthUser) TableName() string {
	return "users"
}

type LoginEvent struct {
	User         *AuthUser
	LoginAllowed *bool
	Error        *ErrorResponse
	Response     *AuthResponse
}

// RegisterRequest represents the payload for user registration
// @Description Registration request payload
type RegisterRequest struct {
	// @Description User's full name
	Name string `json:"name" example:"John Doe" gorm:"column:name"`
	// @Description Username for the account
	Username string `json:"username" example:"johndoe" gorm:"column:username"`
	// @Description User's email address
	Email string `json:"email" binding:"required,email" example:"john@example.com"`
	// @Description Password for the account (minimum 8 characters)
	Password string `json:"password" binding:"required,min=8" example:"password123"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email" example:"john@example.com"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email" example:"john@example.com"`
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6" example:"newpassword123"`
}

type AuthResponse struct {
	users.UserResponse
	AccessToken string      `json:"accessToken"`
	Exp         int64       `json:"exp"`
	Extend      interface{} `json:"extend,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

// VerifyOTPRequest represents the payload to verify an OTP for login
type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required"`
}

// SendOTPRequest represents the payload to request sending an OTP
type SendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}
