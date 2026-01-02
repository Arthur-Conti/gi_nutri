package main

import (
	"github.com/Arthur-Conti/gi_nutri/internal/infra/container"
	"github.com/Arthur-Conti/gi_nutri/internal/infra/http/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	container.BaseContainer = container.NewContainer()
	container.BaseContainer.Start()
	
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
