package route

import (
	"net/http"
	"simple-gateway/api"
	"simple-gateway/pkg"
	"simple-gateway/pkg/gate-way"
	"simple-gateway/service/model"
	routeService "simple-gateway/service/route"
	"simple-gateway/util"
	"time"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary 创建路由
// @Description 创建路由
// @Tags    路由管理
// @Accept  json
// @Produce  json
// @Param body body api.RouteRequest true "请求参数"
// @Success 200 {string} string ""
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /routes [post]
func Create(c *gin.Context) {
	var request = &api.RouteRequest{}
	pkg.ParseRequest(c, request)

	now := time.Now()
	route := model.Route{
		ID:           util.GetRandomString(11),
		Name:         request.Name,
		Path:         request.Path,
		ServerMethod: request.ServerMethod,
		ServerScheme: request.ServerScheme,
		ServiceName:  request.ServiceName,
		ServerHost:   request.ServerHost,
		ServerPort:   request.ServerPort,
		ServerPath:   request.ServerPath,
		GroupID:      request.GroupID,
		Auth:         request.Auth,
		Status:       request.Status,
		Created:      now,
		Updated:      now,
	}
	routeService.Create(&route)

	gate_way.Refresh()
	c.JSON(http.StatusOK, "")
}
