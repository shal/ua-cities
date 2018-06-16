package cmd

import (
	"github.com/ashanaakh/ua-cities/api"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.GET("/city/:name", api.HandleCity)

	// TODO: Make port flexible.
	router.Run(":8080")
}
