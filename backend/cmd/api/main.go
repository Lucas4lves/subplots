package main

import (
	"github.com/Lucas4lves/subplots/database"
	"github.com/Lucas4lves/subplots/internal/container"
	"github.com/Lucas4lves/subplots/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	env := utils.LoadEnv()

	dbCfg := database.NewDatabaseConfig("postgres", env.DSN)

	defer dbCfg.Driver.Close()

	dc := container.NewDependencyContainer(dbCfg.Driver)

	server.Run("0.0.0.0:" + env.PORT)
}
