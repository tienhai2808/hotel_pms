package usecase

import (
	"context"

	"github.com/InstaySystem/is_v2-be/internal/application/dto"
	"github.com/InstaySystem/is_v2-be/internal/domain/model"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, userID int64, req dto.CreateUserRequest) (int64, error)

	GetUserByID(ctx context.Context, userID int64) (*model.User, error)

	GetUsers(ctx context.Context, query dto.UserPaginationQuery) ([]*model.User, *dto.MetaResponse, error)

	UpdateUser(ctx context.Context, userID, currentUserID int64, req dto.UpdateUserRequest) error
}
