package usecase

import (
	"context"

	"github.com/InstayPMS/backend/internal/application/dto"
	"github.com/InstayPMS/backend/internal/domain/model"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, userID int64, req dto.CreateUserRequest) (int64, error)

	GetUserByID(ctx context.Context, userID int64) (*model.User, error)
}