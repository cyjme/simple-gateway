package route

import (
	"net/http"
	"simple-gateway/pkg/gate-way"
	routeService "simple-gateway/service/route"

	"github.com/gin-gonic/gin"
)

// Delete godoc
// @Summary 删除路由
// @Description 删除路由
// @Tags    路由管理
// @Accept  json
// @Produce  json
// @Param string string "" true "请求参数"
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /routes/id [delete]
func Delete(c *gin.Context) {
	routeID := c.Param("id")
	routeService.Delete(routeID)

	gate_way.Refresh()
	c.JSON(http.StatusOK, "")
}
