package router

import (
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/handler"
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/middleware"
	"github.com/gin-gonic/gin"
)

func (r *Router) setupOutletRoutes(rg *gin.RouterGroup, authMid *middleware.AuthMiddleware, hdl *handler.OutletHandler) {
	outlet := rg.Group("/outlets", authMid.IsAuthentication(), authMid.IsSystemAdministrator())
	{
		outlet.POST("", hdl.CreateOutlet)
	}
}