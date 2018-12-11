package route

import (
	"fmt"
	"simple-gateway/api"
	"simple-gateway/app"
	"simple-gateway/service/model"
)

func Create(route *model.Route) {
	app.DB.Create(route)
}

func Update(route *model.Route) {
	fmt.Println("route", route)
	app.DB.Table("route").Where("id = ?", route.ID).Updates(route)
}

func Delete(id string) {
	app.DB.Table("route").Where("id = ?", id).Delete(&model.Route{})
}

func GetByGroupId(groupID string) *[]api.RouteResponse {
	var routes []api.RouteResponse
	app.DB.Table("route").Where("group_id = ?", groupID).Find(&routes)

	return &routes
}

func GetAll() *[]model.Route {
	var routes []model.Route
	app.DB.Table("route").Find(&routes)

	return &routes
}

func CreateRouteGroup(group *model.RouteGroup) {
	app.DB.Create(group)
}

func UpdateRouteGroup(group *model.RouteGroup) {
	app.DB.Table("route_group").Where("id = ?", group.ID).Updates(group)
}

func DeleteRouteGroup(groupID string) {
	app.DB.Table("route_group").Where("id = ?", groupID).Delete(&model.RouteGroup{})
}

func GetRouteGroup() *[]api.RouteGroupResponse {
	//var groups []model.RouteGroup
	var groups []api.RouteGroupResponse
	app.DB.Table("route_group").Find(&groups)

	return &groups
}
