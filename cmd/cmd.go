package cmd

import (
	"github.com/ashanaakh/ua-cities/api"
	"github.com/gin-gonic/gin"
)

// Run executes main project logic
func Run() {
	router := gin.Default()
	api := api.NewCitiesAPI()

	router.GET("/city/:name", api.HandleCity)
	router.GET("/city", api.HandleCityWithQuery)
	router.GET("/cities", api.HandleCities)

	router.Run()
}
