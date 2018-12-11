package route

import (
	"net/http"
	"simple-gateway/api"
	"simple-gateway/pkg"
	"simple-gateway/pkg/gate-way"
	"simple-gateway/service/model"
	routeService "simple-gateway/service/route"
	"time"

	"github.com/gin-gonic/gin"
)

// UpdateRoute godoc
// @Summary 修改路由
// @Description 修改路由
// @Tags    路由管理
// @Accept  json
// @Produce  json
// @Param body body api.RouteRequest true "请求参数"
// @Success 200 {string} string ""
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /routes/id [put]
func Update(c *gin.Context) {
	routeID := c.Param("id")
	request := &api.RouteRequest{}
	pkg.ParseRequest(c, request)

	now := time.Now()

	route := model.Route{
		ID:           routeID,
		Name:         request.Name,
		Path:         request.Path,
		ServerMethod: request.ServerMethod,
		ServerScheme: request.ServerScheme,
		ServerHost:   request.ServerHost,
		ServiceName:  request.ServiceName,
		ServerPort:   request.ServerPort,
		ServerPath:   request.ServerPath,
		GroupID:      request.GroupID,
		Auth:         request.Auth,
		Status:       request.Status,
		Updated:      now,
	}
	routeService.Update(&route)
	gate_way.Refresh()

	c.JSON(http.StatusOK, "")
}
