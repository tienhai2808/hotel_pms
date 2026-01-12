package outlet

import (
	"context"

	"github.com/InstayPMS/backend/internal/application/dto"
)

type OutletUseCase interface {
	CreateOutlet(ctx context.Context, userID int64, req dto.CreateOutletRequest) (int64, error)
}
