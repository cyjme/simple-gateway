package route

import (
	"net/http"
	routeService "simple-gateway/service/route"

	"github.com/gin-gonic/gin"
)

// GetByGroupID godoc
// @Summary     根据 group_id 获取路由
// @Description 根据 group_id 获取路由
// @Tags    路由管理
// @Accept  json
// @Produce  json
// @Success 200 {array} api.RouteResponse
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /routeGroup/id/routes [get]
func GetByGroupID(c *gin.Context) {
	groupID := c.Param("id")
	routes := routeService.GetByGroupId(groupID)
	c.JSON(http.StatusOK, routes)
}
