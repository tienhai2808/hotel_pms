package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/InstayPMS/backend/internal/application/dto"
	userUC "github.com/InstayPMS/backend/internal/application/usecase/user"
	"github.com/InstayPMS/backend/internal/domain/model"
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/middleware"
	"github.com/InstayPMS/backend/pkg/constants"
	"github.com/InstayPMS/backend/pkg/errors"
	"github.com/InstayPMS/backend/pkg/utils"
	"github.com/InstayPMS/backend/pkg/validator"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUC userUC.UserUseCase
}

func NewUserHandler(userUC userUC.UserUseCase) *UserHandler {
	return &UserHandler{userUC}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userID := c.GetInt64(middleware.CtxUserID)
	if userID == 0 {
		c.Error(errors.ErrUnAuth)
		return
	}

	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		field, tag, param := validator.HandleRequestError(err)
		c.Error(errors.ErrBadRequest.WithData(gin.H{
			"field": field,
			"tag":   tag,
			"param": param,
		}))
		return
	}

	if req.Role == model.RoleAdmin {
		if req.DepartmentID != nil {
			c.Error(errors.ErrBadRequest.WithData(gin.H{
				"field": "departmentid",
				"tag":   "notrequired",
				"param": "",
			}))
			return
		}
	} else {
		if req.DepartmentID == nil {
			c.Error(errors.ErrBadRequest.WithData(gin.H{
				"field": "departmentid",
				"tag":   "required",
				"param": "",
			}))
			return
		}
	}

	if req.OutletID == nil && req.DepartmentID != nil {
		c.Error(errors.ErrBadRequest.WithData(gin.H{
			"field": "outletid",
			"tag":   "required",
			"param": "",
		}))
		return
	}

	userID, err := h.userUC.CreateUser(ctx, userID, req)
	if err != nil {
		c.Error(err)
		return
	}

	utils.APIResponse(c, http.StatusCreated, constants.CodeCreateUserSuccess, "User created successfully", gin.H{
		"user_id": userID,
	})
}
