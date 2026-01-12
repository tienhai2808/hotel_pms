package repository

import (
	"context"

	"github.com/InstayPMS/backend/internal/domain/model"
)

type OutletRepository interface {
	Create(ctx context.Context, outlet *model.Outlet) error
}