package gate_way

import (
	"log"
	"simple-gateway/httprouter"
	"simple-gateway/service/model"
	routeService "simple-gateway/service/route"
	"strings"

	"github.com/cyjme/service-center/client"
	"github.com/cyjme/service-center/service"
)

var Router *httprouter.Router
var RouteMap = make(map[string]string)
var BackendRouteMap = make(map[string]model.Route)

type NodeList []string

var ServiceNode = make(map[string]NodeList)

func Init() {
	Router = httprouter.New()
}

func AddRoute(method string, path string, routeId string) {
	Router.Handle(method, path, routeId)
}

func FindRoute(method string, path string) (string, httprouter.Params, bool) {
	routeId, param, tsr := Router.Lookup(method, path)

	return routeId, param, tsr
}

func Refresh() {
	Router.ClearRoute("GET")
	Router.ClearRoute("POST")
	Router.ClearRoute("PUT")
	Router.ClearRoute("DELETE")
	Router.ClearRoute("PATCH")
	Router.ClearRoute("OPTION")
	Router.ClearRoute("HEADER")
	LoadRouteFromMysql()
}

func LoadRouteFromMysql() {
	routes := routeService.GetAll()

	for _, route := range *routes {
		AddRoute(route.ServerMethod, route.Path, route.ID)

		RouteMap[route.ID] = route.Path
		BackendRouteMap[route.ID] = route
	}

	log.Println("LoadRouteFromMysql Success !!!")
}

func LoadAllServiceNode() {
	client.Init("localhost:2379")
	allServiceNode := service.ListNodeByServiceName("")
	ServiceNode = make(map[string]NodeList)

	for _, node := range allServiceNode {
		arr := strings.Split(node.Key, "/")
		serviceName := arr[1]
		ServiceNode[serviceName] = append(ServiceNode[serviceName], node.Value)
	}
	log.Println("Service Node", ServiceNode)
}

func callback() {
	log.Println("watch service change, reloadAllServiceNode")
	LoadAllServiceNode()
}

func WatchService() {
	service.Watch("", callback)
}
