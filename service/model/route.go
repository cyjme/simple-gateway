package model

import "time"

// Route route
type Route struct {
	ID           string
	Name         string
	Path         string
	ServerMethod string
	ServerScheme string
	ServiceName  string
	ServerHost   string
	ServerPort   int64
	ServerPath   string
	GroupID      string
	Auth         bool
	Status       bool
	Created      time.Time
	Updated      time.Time
}

// RouteGroup collection of route
type RouteGroup struct {
	ID      string
	Name    string
	Type    string
	Status  uint
	Created time.Time
	Updated time.Time
}
