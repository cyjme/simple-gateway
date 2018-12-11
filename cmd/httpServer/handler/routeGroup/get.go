package routeGroup

import (
	"github.com/gin-gonic/gin"
	"net/http"
	routeService "simple-gateway/service/route"
)

// GetRouteByGroup godoc
// @Summary    获取路由分组
// @Description 获取路由分组
// @Tags    路由分组管理
// @Accept  json
// @Produce  json
// @Success 200 {array} api.RouteGroupResponse
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /routeGroup [get]
func Get(c *gin.Context) {
	groups := routeService.GetRouteGroup()

	c.JSON(http.StatusOK, groups)
}
