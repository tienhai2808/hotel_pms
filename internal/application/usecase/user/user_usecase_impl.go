package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/InstaySystem/is_v2-be/internal/application/dto"
	"github.com/InstaySystem/is_v2-be/internal/application/port"
	"github.com/InstaySystem/is_v2-be/internal/domain/model"
	"github.com/InstaySystem/is_v2-be/internal/domain/repository"
	customErr "github.com/InstaySystem/is_v2-be/pkg/errors"
	"github.com/InstaySystem/is_v2-be/pkg/utils"
	"github.com/sony/sonyflake/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userUseCaseImpl struct {
	db        *gorm.DB
	log       *zap.Logger
	idGen     *sonyflake.Sonyflake
	cachePro  port.CacheProvider
	userRepo  repository.UserRepository
	deptRepo  repository.DepartmentRepository
	tokenRepo repository.TokenRepository
}

func NewUserUseCase(
	db *gorm.DB,
	log *zap.Logger,
	idGen *sonyflake.Sonyflake,
	cachePro port.CacheProvider,
	userRepo repository.UserRepository,
	deptRepo repository.DepartmentRepository,
	tokenRepo repository.TokenRepository,
) UserUseCase {
	return &userUseCaseImpl{
		db,
		log,
		idGen,
		cachePro,
		userRepo,
		deptRepo,
		tokenRepo,
	}
}

func (u *userUseCaseImpl) CreateUser(ctx context.Context, userID int64, req dto.CreateUserRequest) (int64, error) {
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
		IsActive:     *req.IsActive,
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
			return 0, customErr.ErrDepartmentNotFound
		}
		u.log.Error("create user failed", zap.Error(err))
		return 0, err
	}

	return id, nil
}

func (u *userUseCaseImpl) GetUserByID(ctx context.Context, userID int64) (*model.User, error) {
	user, err := u.userRepo.FindByIDWithDetails(ctx, userID)
	if err != nil {
		u.log.Error("find user by id failed", zap.Int64("id", userID), zap.Error(err))
		return nil, err
	}
	if user == nil {
		return nil, customErr.ErrUserNotFound
	}

	return user, nil
}

func (u *userUseCaseImpl) GetUsers(ctx context.Context, query dto.UserPaginationQuery) ([]*model.User, *dto.MetaResponse, error) {
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Limit == 0 {
		query.Limit = 10
	}

	users, total, err := u.userRepo.FindAllWithDepartmentPaginated(ctx, query)
	if err != nil {
		u.log.Error("find all users paginated failed", zap.Error(err))
		return nil, nil, err
	}

	meta := utils.CalculateMeta(total, query.Page, query.Limit)

	return users, meta, nil
}

func (u *userUseCaseImpl) UpdateUser(ctx context.Context, userID, currentUserID int64, req dto.UpdateUserRequest) error {
	if userID == currentUserID && (*req.IsActive == false || req.Role == model.RoleStaff) {
		exists, err := u.userRepo.ExistsActiveAdminExceptID(ctx, userID)
		if err != nil {
			u.log.Error("check active admin except id failed", zap.Int64("id", userID), zap.Error(err))
			return err
		}
		if !exists {
			return customErr.ErrNeedAdmin
		}
	}

	updateData := map[string]any{
		"username":      req.Username,
		"email":         req.Username,
		"phone":         req.Phone,
		"first_name":    req.FirstName,
		"last_name":     req.LastName,
		"role":          req.Role,
		"is_active":     *req.IsActive,
		"department_id": req.DepartmentID,
		"updated_by_id": currentUserID,
	}

	if err := u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := u.userRepo.UpdateTx(tx, userID, updateData); err != nil {
			if errors.Is(err, customErr.ErrUserNotFound) {
				return err
			}
			if ok, constraint := utils.IsUniqueViolation(err); ok {
				switch constraint {
				case "users_email_key":
					return customErr.ErrEmailAlreadyExists
				case "users_username_key":
					return customErr.ErrUsernameAlreadyExists
				case "users_phone_key":
					return customErr.ErrPhoneAlreadyExists
				}
			}
			if ok, _ := utils.IsForeignKeyViolation(err); ok {
				return customErr.ErrDepartmentNotFound
			}
			u.log.Error("update user failed", zap.Int64("id", userID), zap.Error(err))
			return err
		}

		if *req.IsActive == false {
			if err := u.tokenRepo.UpdateByUserIDTx(tx, userID, map[string]any{"revoked_at": time.Now()}); err != nil {
				u.log.Error("update token by user id failed", zap.Error(err))
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	if *req.IsActive == false {
		redisKey := fmt.Sprintf("user_version:%d", userID)
		if err := u.cachePro.Increment(ctx, redisKey); err != nil {
			u.log.Error("increase token version failed", zap.Error(err))
		}
	}

	return nil
}
