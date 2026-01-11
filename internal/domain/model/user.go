package model

import "time"

type UserRole string

const (
	RoleStaff UserRole = "staff"
	RoleAdmin UserRole = "admin"
)

type User struct {
	ID           int64     `gorm:"type:bigint;primaryKey" json:"id"`
	Username     string    `gorm:"type:varchar(50);uniqueIndex:users_username_key;not null" json:"username"`
	Email        string    `gorm:"type:varchar(150);uniqueIndex:users_email_key;not null" json:"email"`
	Role         UserRole  `gorm:"type:varchar(20);check:role IN ('staff', 'admin')" json:"role"`
	FirstName    string    `gorm:"type:varchar(150);not null" json:"first_name"`
	LastName     string    `gorm:"type:varchar(150);not null" json:"last_name"`
	Phone        string    `gorm:"type:char(10);uniqueIndex:users_phone_key;not null" json:"phone"`
	Password     string    `gorm:"type:varchar(255);not null" json:"password"`
	IsActive     bool      `gorm:"type:boolean;not null" json:"is_active"`
	OutletID     *int64    `gorm:"type:bigint" json:"outlet_id"`
	DepartmentID *int64    `gorm:"type:bigint" json:"department_id"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedByID  *int64    `gorm:"type:bigint" json:"created_by_id"`
	UpdatedByID  *int64    `gorm:"type:bigint" json:"updated_by_id"`

	Outlet             *Outlet       `gorm:"foreignKey:OutletID;references:ID;OnUpdate:CASCADE,OnDelete:RESTRICT" json:"outlet"`
	Department         *Department   `gorm:"foreignKey:DepartmentID;references:ID;OnUpdate:CASCADE,OnDelete:RESTRICT" json:"department"`
	CreatedBy          *User         `gorm:"foreignKey:CreatedByID;references:ID;OnUpdate:CASCADE,OnDelete:RESTRICT" json:"created_by"`
	UpdatedBy          *User         `gorm:"foreignKey:UpdatedByID;references:ID;OnUpdate:CASCADE,OnDelete:RESTRICT" json:"updated_by"`
	Tokens             []*Token      `gorm:"foreignKey:UserID;references:ID" json:"tokens"`
}

func IsValidRole(str string) bool {
	switch str {
	case string(RoleAdmin):
		return true
	case string(RoleStaff):
		return true
	default:
		return false
	}
}
