package repository

import (
	"context"

	"github.com/InstayPMS/backend/internal/domain/model"
)

type TokenRepository interface {
	Create(ctx context.Context, token *model.Token) error
}