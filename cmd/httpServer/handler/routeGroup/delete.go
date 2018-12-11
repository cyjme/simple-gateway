package routeGroup

import (
	"net/http"
	routeService "simple-gateway/service/route"

	"github.com/gin-gonic/gin"
)

// DeleteRouteGroup godoc
// @Summary 删除路由分组
// @Description 创建路由分组
// @Tags    路由分组管理
// @Accept  json
// @Produce  json
// @Success 200 {string} string ""
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /routeGroup/:id [delete]
func Delete(c *gin.Context) {
	groupID := c.Param("id")
	routeService.DeleteRouteGroup(groupID)

	c.JSON(http.StatusOK, "")
}
