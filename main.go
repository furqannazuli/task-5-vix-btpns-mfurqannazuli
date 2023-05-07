package main

import (
	"github.com/furqannazuli/task-5-vix-btpns-mfurqannazuli/database"
	"github.com/furqannazuli/task-5-vix-btpns-mfurqannazuli/router"
)

func main() {
	database.InitDB()
	database.MigrateDB()
	r := router.RouteInit()
	r.Run(":8080")
}
