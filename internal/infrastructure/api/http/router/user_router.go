package router

import (
	"github.com/InstaySystem/is_v2-be/internal/domain/model"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/api/http/handler"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/api/http/middleware"
	"github.com/gin-gonic/gin"
)

func (r *Router) setupUserRoutes(rg *gin.RouterGroup, authMid *middleware.AuthMiddleware, hdl *handler.UserHandler) {
	user := rg.Group("/users", authMid.IsAuthentication(), authMid.HasRole(model.RoleAdmin))
	{
		user.POST("", hdl.CreateUser)

		user.GET("/:id", hdl.GetUserByID)

		user.GET("", hdl.GetUsers)

		user.GET("/roles", hdl.GetAllRoles)

		user.PUT("/:id", hdl.UpdateUser)
	}
}
