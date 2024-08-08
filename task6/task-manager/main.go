package main

import (
	"task-manager/routers"
	"task-manager/data"
)

func main() {
	data.InitDB()
	routers.SetupRouter().Run(":8080")
}