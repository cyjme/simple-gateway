package httpServer

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"simple-gateway/app"
	"simple-gateway/cmd/httpServer/middleware"
	"simple-gateway/cmd/httpServer/router"
	_ "simple-gateway/docs"
	"simple-gateway/pkg/gate-way"
	"simple-gateway/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"math/rand"
	"strconv"
)

func StartServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if gin.Mode() == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Use(middleware.CORSMiddleware())
	router.RegisterV1Router(r)

	r.Any("/api/*path", ReverseProxy())

	r.Run(app.Config.Http.Domain + ":" + app.Config.Http.Port)
}

func ReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path)
		method := c.Request.Method
		path := c.Request.URL.Path
		path = strings.TrimLeft(path, "/")
		path = strings.TrimLeft(path, "api")
		routeID, param, _ := gate_way.FindRoute(method, path)
		if routeID == "" {
			c.JSON(http.StatusBadGateway, "api gateWay: route not register")
			return
		}
		backendRoute := gate_way.BackendRouteMap[routeID]
		if backendRoute.Auth {
			userID, err := parseToken(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, "api gateWay: route required auth token")
				return
			}
			c.Request.Header.Set("user_id", userID)
		}
		backendRoutePath := backendRoute.ServerPath
		for _, pathParam := range param {
			backendRoutePath = strings.Replace(backendRoutePath, ":"+pathParam.Key, pathParam.Value, 1)
		}
		factoryID := c.GetHeader("Factory_id")
		log.Println("factory_id in header", factoryID)
		if factoryID != "" {
			c.Request.Header.Set("factory_id", factoryID)
		}

		req := c.Request

		//如果 serviceName 存在，根据serviceName 从服务中心获取对应的真实 ip
		serviceLen := len(gate_way.ServiceNode[backendRoute.ServiceName])
		log.Println("serviceName",backendRoute.ServiceName)
		log.Println("serviceLen",serviceLen)
		if backendRoute.ServiceName != "" && serviceLen > 0 {
			req.URL.Host = gate_way.ServiceNode[backendRoute.ServiceName][rand.Intn(serviceLen)]
			log.Println("reverse to ",req.URL.Host)
		} else {
			req.URL.Host = backendRoute.ServerHost + ":" + strconv.Itoa(int(backendRoute.ServerPort))
		}

		req.URL.Scheme = backendRoute.ServerScheme
		req.URL.Path = backendRoutePath
		director := func(req *http.Request) {
		}
		log.Println("reverse url", req.URL)
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func parseToken(c *gin.Context) (string, error) {
	authorization := c.Request.Header.Get("Authorization")

	if !strings.Contains(authorization, "Bearer") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return "", errors.New("no Bearer")
	}

	token := string([]byte(authorization)[7:])

	if token == "null" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return "", errors.New("token is null")
	}

	mc, err := util.ParseJwtToken(app.Config.Secret.JwtKey, token)
	if err != nil {
		return "", err
	}

	return mc["user_id"].(string), nil
}
