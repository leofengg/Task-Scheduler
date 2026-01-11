package main

import (
	"fmt"

	"./src/scheduler"

	"github.com/JyotinderSingh/task-queue/pkg/common"
)

func main() {
	dbConnectionString := common.GetDBConnectionString()
	schedulerServer := scheduler.NewServer("8001", dbConnectionString)
	schedulerServer.StartServer()
	fmt.Println("Starting scheduler")
}
