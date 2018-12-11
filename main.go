package main

import (
	"simple-gateway/app"
	"simple-gateway/cmd/httpServer"
	"simple-gateway/pkg/gate-way"
	"fmt"
)

func main() {
	fmt.Println("simple-gateway starting...")
	app.InitConfig()
	app.InitDB()
	gate_way.Init()
	gate_way.LoadRouteFromMysql()
	gate_way.LoadAllServiceNode()
	go gate_way.WatchService()
	httpServer.StartServer()
}
