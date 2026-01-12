package repository

import (
	"context"

	"github.com/InstaySystem/is_v2-be/internal/domain/model"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Create(ctx context.Context, token *model.Token) error

	UpdateByToken(ctx context.Context, token string, updateData map[string]any) error

	FindByToken(ctx context.Context, token string) (*model.Token, error)

	UpdateByUserIDTx(tx *gorm.DB, userID int64, updateData map[string]any) error
}
