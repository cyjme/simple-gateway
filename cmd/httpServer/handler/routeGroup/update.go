package routeGroup

import (
	"net/http"
	"simple-gateway/api"
	"simple-gateway/service/model"
	routeService "simple-gateway/service/route"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateRoute godoc
// @Summary 修改路由分组
// @Description 修改路由分组
// @Tags    路由分组管理
// @Accept  json
// @Produce  json
// @Param body body api.RouteGroupRequest true "请求参数"
// @Success 200 {string} string ""
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /routeGroup/:id [put]
func Update(c *gin.Context) {
	id := c.Param("id")
	request := &api.RouteGroupRequest{}
	if err := c.ShouldBindWith(request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	now := time.Now()
	group := model.RouteGroup{
		ID:      id,
		Name:    request.Name,
		Type:    request.Type,
		Status:  0,
		Updated: now,
	}
	routeService.UpdateRouteGroup(&group)

	c.JSON(http.StatusOK, "")
}
