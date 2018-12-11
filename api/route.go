package api

import "time"

//RouteRequest create or update route request
type RouteRequest struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	ServerMethod string `json:"serverMethod"`
	ServerScheme string `json:"serverScheme"`
	ServiceName  string `json:"serviceName"`
	ServerHost   string `json:"serverHost"`
	ServerPort   int64  `json:"serverPort"`
	ServerPath   string `json:"serverPath"`
	GroupID      string `json:"groupId"`
	Auth         bool   `json:"auth"`
	Status       bool   `json:"status"`
}

//RouteResponse fetch Route response
type RouteResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	ServerMethod string    `json:"serverMethod"`
	ServerScheme string    `json:"serverScheme"`
	ServiceName   string    `json:"serviceName"`
	ServerHost   string    `json:"serverHost"`
	ServerPort   int64     `json:"serverPort"`
	ServerPath   string    `json:"serverPath"`
	GroupID      string    `json:"groupId"`
	Auth         bool      `json:"auth"`
	Status       bool      `json:"status"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
}

//RouteGroupRequest create or update routeGroup
type RouteGroupRequest struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

//RouteGroupResponse fetch RouteGroup Response
type RouteGroupResponse struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Type    string    `json:"type"`
	Status  string    `json:"status"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
