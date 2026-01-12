package dto

import "github.com/InstaySystem/is_v2-be/internal/domain/model"

type UploadPresignedURLRequest struct {
	FileName    string `json:"file_name" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
}

type UploadPresignedURLsRequest struct {
	Files []UploadPresignedURLRequest `json:"files" binding:"required,min=1,dive"`
}

type ViewPresignedURLsRequest struct {
	Keys []string `json:"keys" binding:"required,min=1,dive"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=6"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=6"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyForgotPasswordRequest struct {
	ForgotPasswordToken string `json:"forgot_password_token" binding:"required,uuid4"`
	Otp                 string `json:"otp" binding:"required,len=6,numeric"`
}

type ResetPasswordRequest struct {
	ResetPasswordToken string `json:"reset_password_token" binding:"required,uuid4"`
	NewPassword        string `json:"new_password" binding:"required,min=6"`
}

type UpdateInfoRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,len=10"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type CreateUserRequest struct {
	Username     string         `json:"username" binding:"required,min=5"`
	Email        string         `json:"email" binding:"required,email"`
	Phone        string         `json:"phone" binding:"required,len=10"`
	Password     string         `json:"password" binding:"required,min=6"`
	Role         model.UserRole `json:"role" binding:"required,oneof=staff admin"`
	IsActive     *bool          `json:"is_active" binding:"required"`
	FirstName    string         `json:"first_name" binding:"required"`
	LastName     string         `json:"last_name" binding:"required"`
	DepartmentID *int64         `json:"department_id" binding:"omitempty"`
}

type CreateDepartmentRequest struct {
	Name        string `json:"name" binding:"required,min=2"`
	Phone       string `json:"phone" binding:"required,max=20"`
	Description string `json:"description" binding:"required,min=1"`
	IsActive    bool   `json:"is_active" binding:"required"`
}

type UserPaginationQuery struct {
	Page         uint32 `form:"page" binding:"omitempty,min=1" json:"page"`
	Limit        uint32 `form:"limit" binding:"omitempty,min=1,max=100" json:"limit"`
	Sort         string `form:"sort" json:"sort"`
	Order        string `form:"order" binding:"omitempty,oneof=asc desc" json:"order"`
	Role         string `form:"role" binding:"omitempty,oneof=admin staff" json:"role"`
	DepartmentID int64  `form:"department_id" binding:"omitempty" json:"department_id"`
	IsActive     *bool  `form:"is_active" binding:"omitempty" json:"is_active"`
	Search       string `form:"search" json:"search"`
}

type UpdateUserRequest struct {
	Username     string         `json:"username" binding:"required,min=5"`
	Email        string         `json:"email" binding:"required,email"`
	Phone        string         `json:"phone" binding:"required,len=10"`
	FirstName    string         `json:"first_name" binding:"required"`
	LastName     string         `json:"last_name" binding:"required"`
	Role         model.UserRole `json:"role" binding:"required,oneof=staff admin"`
	IsActive     *bool          `json:"is_active" binding:"required"`
	DepartmentID *int64         `json:"department_id" binding:"omitempty"`
}
