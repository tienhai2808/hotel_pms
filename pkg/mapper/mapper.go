package mapper

import (
	"github.com/InstayPMS/backend/internal/application/dto"
	"github.com/InstayPMS/backend/internal/domain/model"
)

func ToSimpleOutletResponse(ol *model.Outlet) *dto.SimpleOutletResponse {
	if ol == nil {
		return nil
	}

	return &dto.SimpleOutletResponse{
		ID:   ol.ID,
		Name: ol.Name,
	}
}

func ToBasicDepartmentResponse(de *model.Department) *dto.BasicDepartmentResponse {
	if de == nil {
		return nil
	}

	return &dto.BasicDepartmentResponse{
		ID:   de.ID,
		Name: de.Name,
	}
}

func ToUserResponse(us *model.User) *dto.UserResponse {
	if us == nil {
		return nil
	}

	return &dto.UserResponse{
		ID:         us.ID,
		Email:      us.Email,
		Phone:      us.Phone,
		Username:   us.Username,
		FirstName:  us.FirstName,
		LastName:   us.LastName,
		Role:       us.Role,
		IsActive:   us.IsActive,
		CreatedAt:  us.CreatedAt,
		Outlet:     ToSimpleOutletResponse(us.Outlet),
		Department: ToBasicDepartmentResponse(us.Department),
	}
}
