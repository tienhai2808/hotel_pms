package orm

import (
	"context"

	"github.com/InstayPMS/backend/internal/domain/model"
	"github.com/InstayPMS/backend/internal/domain/repository"
	"gorm.io/gorm"
)

type outletRepositoryImpl struct {
	db *gorm.DB
}

func NewOutletRepository(db *gorm.DB) repository.OutletRepository {
	return &outletRepositoryImpl{db}
}

func (r *outletRepositoryImpl) Create(ctx context.Context, outlet *model.Outlet) error {
	return r.db.WithContext(ctx).Create(outlet).Error
}
