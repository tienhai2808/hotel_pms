package model

import "time"

type Token struct {
	ID        int64      `gorm:"type:bigint;primaryKey" json:"id"`
	UserID    int64      `gorm:"type:bigint;not null" json:"user_id"`
	Token     string     `gorm:"type:string;not null;uniqueIndex:tokens_token_key" json:"token"`
	UserAgent string     `gorm:"type:string;not null" json:"user_agent"`
	IPAddress string     `gorm:"type:string;not null" json:"id_address"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	RevokedAt *time.Time `json:"revoked_at"`
	ExpiresAt time.Time  `gorm:"not null" json:"expires_at"`

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:fk_tokens_user,OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
}
