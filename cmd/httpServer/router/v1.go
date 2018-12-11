package router

import (
	"simple-gateway/cmd/httpServer/handler/route"
	"simple-gateway/cmd/httpServer/handler/routeGroup"

	"github.com/gin-gonic/gin"
)

func RegisterV1Router(r *gin.Engine) {
	r.POST("/routes", route.Create)
	r.DELETE("/routes/:id", route.Delete)
	r.PUT("/routes/:id", route.Update)
	r.GET("/routeGroup/:id/routes", route.GetByGroupID)

	r.POST("/routeGroup", routeGroup.Create)
	r.DELETE("/routeGroup/:id", routeGroup.Delete)
	r.PUT("/routeGroup/:id", routeGroup.Update)
	r.GET("/routeGroup", routeGroup.Get)
}
