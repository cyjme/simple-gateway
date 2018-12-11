package routeGroup

import (
	"net/http"
	"simple-gateway/api"
	"simple-gateway/pkg"
	"simple-gateway/service/model"
	routeService "simple-gateway/service/route"
	"simple-gateway/util"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateRouteGroup godoc
// @Summary 创建路由分组
// @Description 创建路由分组
// @Tags    路由分组管理
// @Accept  json
// @Produce  json
// @Param body body api.RouteGroupRequest true "请求参数"
// @Success 200 {string} string ""
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /routeGroup [post]
func Create(c *gin.Context) {
	request := &api.RouteGroupRequest{}
	pkg.ParseRequest(c, request)

	now := time.Now()
	group := model.RouteGroup{
		ID:      util.GetRandomString(11),
		Type:    request.Type,
		Name:    request.Name,
		Status:  0,
		Created: now,
		Updated: now,
	}
	routeService.CreateRouteGroup(&group)

	c.JSON(http.StatusOK, "")
}
