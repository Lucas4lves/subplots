package main

import (
	"log"

	"github.com/Lucas4lves/subplots/database"
	"github.com/Lucas4lves/subplots/internal/container"
	"github.com/Lucas4lves/subplots/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	env := utils.LoadEnv()

	dbCfg := database.NewDatabaseConfig("postgres", env.DSN)

	err := dbCfg.CreatePlotsTable()

	if err != nil {
		log.Println(err)
	}

	defer dbCfg.Driver.Close()

	dc := container.NewDependencyContainer(dbCfg.Driver)

	server.GET("/plots", dc.PlotController.GetAll)
	server.POST("/plots", dc.PlotController.CreatePlot)

	server.Run("0.0.0.0:" + env.PORT)
}
