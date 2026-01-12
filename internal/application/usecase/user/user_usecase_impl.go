package usecase

import (
	"context"

	"github.com/InstayPMS/backend/internal/application/dto"
	"github.com/InstayPMS/backend/internal/domain/model"
	"github.com/InstayPMS/backend/internal/domain/repository"
	customErr "github.com/InstayPMS/backend/pkg/errors"
	"github.com/InstayPMS/backend/pkg/utils"
	"github.com/sony/sonyflake/v2"
	"go.uber.org/zap"
)

type userUseCaseImpl struct {
	log      *zap.Logger
	idGen    *sonyflake.Sonyflake
	userRepo repository.UserRepository
	deptRepo repository.DepartmentRepository
}

func NewUserUseCase(
	log *zap.Logger,
	idGen *sonyflake.Sonyflake,
	userRepo repository.UserRepository,
	deptRepo repository.DepartmentRepository,
) UserUseCase {
	return &userUseCaseImpl{
		log,
		idGen,
		userRepo,
		deptRepo,
	}
}

func (u *userUseCaseImpl) CreateUser(ctx context.Context, userID int64, req dto.CreateUserRequest) (int64, error) {
	if req.DepartmentID != nil {
		dept, err := u.deptRepo.FindByID(ctx, *req.DepartmentID)
		if err != nil {
			u.log.Error("find department by id failed", zap.Error(err))
			return 0, err
		}
		if dept == nil {
			return 0, customErr.ErrDepartmentNotFound
		}

		if dept.OutletID != *req.OutletID {
			return 0, customErr.ErrDepartmentDoesNotExistInOutlet
		}
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		u.log.Error("hash password failed", zap.Error(err))
		return 0, err
	}

	id, err := u.idGen.NextID()
	if err != nil {
		u.log.Error("generate user id failed", zap.Error(err))
		return 0, err
	}

	user := &model.User{
		ID:           id,
		Username:     req.Username,
		Email:        req.Email,
		Password:     hashedPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Phone:        req.Phone,
		Role:         req.Role,
		IsActive:     req.IsActive,
		OutletID:     req.OutletID,
		DepartmentID: req.DepartmentID,
		CreatedByID:  &userID,
		UpdatedByID:  &userID,
	}

	if err = u.userRepo.Create(ctx, user); err != nil {
		if ok, constraint := utils.IsUniqueViolation(err); ok {
			switch constraint {
			case "users_email_key":
				return 0, customErr.ErrEmailAlreadyExists
			case "users_username_key":
				return 0, customErr.ErrUsernameAlreadyExists
			case "users_phone_key":
				return 0, customErr.ErrPhoneAlreadyExists
			}
		}
		if ok, _ := utils.IsForeignKeyViolation(err); ok {
			return 0, customErr.ErrOutletNotFound
		}
		u.log.Error("create user failed", zap.Error(err))
		return 0, err
	}

	return id, nil
}
