package main

import (
	"github.com/Lucas4lves/subplots/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	env := utils.LoadEnv()

	server.Run("0.0.0.0:" + env.PORT)
}
