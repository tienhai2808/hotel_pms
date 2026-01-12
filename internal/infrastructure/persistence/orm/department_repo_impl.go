package orm

import (
	"context"
	"errors"

	"github.com/InstaySystem/is_v2-be/internal/domain/model"
	"github.com/InstaySystem/is_v2-be/internal/domain/repository"
	"gorm.io/gorm"
)

type departmentRepositoryImpl struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) repository.DepartmentRepository {
	return &departmentRepositoryImpl{db}
}

func (r *departmentRepositoryImpl) Create(ctx context.Context, dept *model.Department) error {
	return r.db.WithContext(ctx).Create(dept).Error
}

func (r *departmentRepositoryImpl) FindByID(ctx context.Context, id int64) (*model.Department, error) {
	var dept model.Department
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&dept).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &dept, nil
}
