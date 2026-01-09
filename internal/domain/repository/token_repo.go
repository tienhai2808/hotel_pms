package repository

import (
	"context"

	"github.com/InstayPMS/backend/internal/domain/model"
)

type TokenRepository interface {
	Create(ctx context.Context, token *model.Token) error

	UpdateByUserIDAndToken(ctx context.Context, userID int64, token string, updateData map[string]any) error

	FindByToken(ctx context.Context, token string) (*model.Token, error)
}
