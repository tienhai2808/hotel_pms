package mapper

import (
	"github.com/InstayPMS/backend/internal/application/dto"
	"github.com/InstayPMS/backend/internal/domain/model"
)

func ToSimpleOutletResponse(outlet *model.Outlet) *dto.SimpleOutletResponse {
	if outlet == nil {
		return nil
	}

	return &dto.SimpleOutletResponse{
		ID:   outlet.ID,
		Name: outlet.Name,
	}
}

func ToBasicDepartmentResponse(dept *model.Department) *dto.BasicDepartmentResponse {
	if dept == nil {
		return nil
	}

	return &dto.BasicDepartmentResponse{
		ID:   dept.ID,
		Name: dept.Name,
	}
}

func ToUserResponse(usr *model.User) *dto.UserResponse {
	if usr == nil {
		return nil
	}

	return &dto.UserResponse{
		ID:         usr.ID,
		Email:      usr.Email,
		Phone:      usr.Phone,
		Username:   usr.Username,
		FirstName:  usr.FirstName,
		LastName:   usr.LastName,
		Role:       usr.Role,
		IsActive:   usr.IsActive,
		CreatedAt:  usr.CreatedAt,
		Outlet:     ToSimpleOutletResponse(usr.Outlet),
		Department: ToBasicDepartmentResponse(usr.Department),
	}
}
