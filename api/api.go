package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	citiesFilePath = "resources/cities.json"
)

type (
	location struct {
		Latitude  string `json:"lat"`
		Longitude string `json:"lon"`
	}

	// City struct uses for city as JSON.
	City struct {
		Name     string   `json:"name"`
		Location location `json:"location"`
	}

	// CitiesAPI struct uses for API.
	CitiesAPI struct {
		cities []City
		path   string
	}
)

// getCity is method that returns pointer to city object with specified name.
func (api *CitiesAPI) getCity(name string) (*City, error) {
	for _, city := range api.cities {
		if strings.EqualFold(city.Name, name) {
			return &city, nil
		}
	}
	return nil, errors.New("invalid city name")
}

// HandleCity is a gin router handler for getting city.
func (api *CitiesAPI) HandleCity(c *gin.Context) {
	name := c.Param("name")
	city, err := api.getCity(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, city)
	}
}

// HandleCityWithQuery is a gin router handler for getting full list of cities.
func (api *CitiesAPI) HandleCityWithQuery(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		id := rand.Intn(len(api.cities))
		c.JSON(http.StatusOK, api.cities[id])
		return
	}

	city, err := api.getCity(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, city)
	}
}

// HandleCities is a gin router handler for getting full list of cities.
func (api *CitiesAPI) HandleCities(c *gin.Context) {
	c.JSON(http.StatusOK, api.cities)
}

// NewCitiesAPI creats new object of API with full list of cities.
func NewCitiesAPI() *CitiesAPI {
	raw, err := ioutil.ReadFile(citiesFilePath)

	if err != nil {
		fmt.Print(err.Error())
	}

	cities := []City{}
	json.Unmarshal(raw, &cities)

	newObj := new(CitiesAPI)
	newObj.cities = cities

	return newObj
}
