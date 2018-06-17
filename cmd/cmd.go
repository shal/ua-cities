package cmd

import (
	"os"

	"github.com/ashanaakh/ua-cities/api"
	"github.com/gin-gonic/gin"
)

// Run executes main project logic
func Run() {
	router := gin.Default()
	api := api.NewCitiesAPI()

	router.GET("/city/:name", api.HandleCity)

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	router.Run(":" + port)
}
