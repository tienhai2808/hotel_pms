package router

import (
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/handler"
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/middleware"
	"github.com/gin-gonic/gin"
)

func (r *Router) setupUserRoutes(rg *gin.RouterGroup, authMid *middleware.AuthMiddleware, hdl *handler.UserHandler) {
	user := rg.Group("/users", authMid.IsAuthentication(), authMid.IsSystemAdministrator())
	{
		user.POST("", hdl.CreateUser)

		user.GET("/:id", hdl.GetUserByID)
	}
}