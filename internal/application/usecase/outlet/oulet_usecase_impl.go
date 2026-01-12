package outlet

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

type outletUseCaseImpl struct {
	log        *zap.Logger
	idGen      *sonyflake.Sonyflake
	outletRepo repository.OutletRepository
}

func NewOutletUseCase(
	log *zap.Logger,
	idGen *sonyflake.Sonyflake,
	outletRepo repository.OutletRepository,
) OutletUseCase {
	return &outletUseCaseImpl{
		log,
		idGen,
		outletRepo,
	}
}

func (u *outletUseCaseImpl) CreateOutlet(ctx context.Context, userID int64, req dto.CreateOutletRequest) (int64, error) {
	id, err := u.idGen.NextID()
	if err != nil {
		u.log.Error("generate outlet id failed", zap.Error(err))
		return 0, err
	}

	outlet := &model.Outlet{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	if err = u.outletRepo.Create(ctx, outlet); err != nil {
		if ok, constraint := utils.IsUniqueViolation(err); ok {
			switch constraint {
			case "outlets_name_key":
				return 0, customErr.ErrNameAlreadyExists
			case "outlets_phone_key":
				return 0, customErr.ErrPhoneAlreadyExists
			}
		}
		u.log.Error("create outlet failed", zap.Error(err))
		return 0, err
	}

	return id, nil
}
