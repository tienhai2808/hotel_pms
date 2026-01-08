package orm

import (
	"context"

	"github.com/InstayPMS/backend/internal/domain/model"
	"github.com/InstayPMS/backend/internal/domain/repository"
	"gorm.io/gorm"
)

type tokenRepositoryImpl struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) repository.TokenRepository {
	return &tokenRepositoryImpl{db}
}

func (r *tokenRepositoryImpl) Create(ctx context.Context, token *model.Token) error {
	return r.db.WithContext(ctx).Create(token).Error
}
