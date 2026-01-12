package orm

import (
	"context"
	"errors"

	"github.com/InstaySystem/is_v2-be/internal/domain/model"
	"github.com/InstaySystem/is_v2-be/internal/domain/repository"
	customErr "github.com/InstaySystem/is_v2-be/pkg/errors"
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

func (r *tokenRepositoryImpl) UpdateByToken(ctx context.Context, token string, updateData map[string]any) error {
	result := r.db.WithContext(ctx).
		Model(&model.Token{}).
		Where("token = ?", token).
		Updates(updateData)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return customErr.ErrInvalidUser
	}

	return nil
}

func (r *tokenRepositoryImpl) UpdateByUserIDTx(tx *gorm.DB, userID int64, updateData map[string]any) error {
	return tx.Model(&model.Token{}).
		Where("user_id = ?", userID).
		Updates(updateData).Error
}

func (r *tokenRepositoryImpl) FindByToken(ctx context.Context, hashedToken string) (*model.Token, error) {
	var token model.Token
	if err := r.db.WithContext(ctx).Where("token = ?", hashedToken).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &token, nil
}
