package model

import "time"

type Department struct {
	ID          int64     `gorm:"type:bigint;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(150);not null;uniqueIndex:departments_name_outlet_id_key" json:"name"`
	Description string    `gorm:"type:text;not null" json:"description"`
	IsActive    bool      `gorm:"type:boolean;not null" json:"is_active"`
	OutletID    int64     `gorm:"type:bigint;not null;uniqueIndex:departments_name_outlet_id_key" json:"outlet_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedByID *int64    `gorm:"type:bigint" json:"created_by_id"`
	UpdatedByID *int64    `gorm:"type:bigint" json:"updated_by_id"`

	Outlet    *Outlet `gorm:"foreignKey:OutletID;references:ID;constraint:fk_departments_outlet,OnUpdate:CASCADE,OnDelete:RESTRICT" json:"outlet"`
	CreatedBy *User   `gorm:"foreignKey:CreatedByID;references:ID;constraint:-" json:"created_by"`
	UpdatedBy *User   `gorm:"foreignKey:UpdatedByID;references:ID;constraint:-" json:"updated_by"`
	Users     []*User `gorm:"foreignKey:DepartmentID;references:ID;constraint:fk_users_department,OnUpdate:CASCADE,OnDelete:RESTRICT" json:"users"`
}
