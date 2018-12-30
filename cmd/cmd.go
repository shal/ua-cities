package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/shal/ua-cities/api"
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
