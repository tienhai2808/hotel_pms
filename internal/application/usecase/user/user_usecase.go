package usecase

import (
	"context"

	"github.com/InstayPMS/backend/internal/application/dto"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, userID int64, req dto.CreateUserRequest) (int64, error)
}