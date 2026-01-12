package router

import (
	"github.com/InstaySystem/is_v2-be/internal/domain/model"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/api/http/handler"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/api/http/middleware"
	"github.com/gin-gonic/gin"
)

func (r *Router) setupDepartmentRoutes(rg *gin.RouterGroup, authMid *middleware.AuthMiddleware, hdl *handler.DepartmentHandler) {
	dept := rg.Group("/departments", authMid.IsAuthentication(), authMid.HasRole(model.RoleAdmin))
	{
		dept.POST("", hdl.CreateDepartment)
	}
}
