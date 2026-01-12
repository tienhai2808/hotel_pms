package department

import (
	"context"

	"github.com/InstaySystem/is_v2-be/internal/application/dto"
)

type DepartmentUseCase interface {
	CreateDepartment(ctx context.Context, userID int64, req dto.CreateDepartmentRequest) (int64, error)
}
