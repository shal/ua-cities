package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type location struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
}

// City is structure for storing city.
type City struct {
	Name     string   `json:"name"`
	Location location `json:"location"`
}

type citiesAPI struct {
	cities []City
}

func (api *citiesAPI) getCity(name string) (*City, error) {
	for _, city := range api.cities {
		if strings.EqualFold(city.Name, name) {
			return &city, nil
		}
	}
	return nil, errors.New("invalid city name")
}

func (api *citiesAPI) HandleCity(c *gin.Context) {
	name := c.Param("name")
	city, err := api.getCity(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
	} else {
		fmt.Print(city)
		c.JSON(http.StatusOK, city)
	}
}

func NewCitiesAPI() *citiesAPI {
	raw, err := ioutil.ReadFile("resources/cities.json")

	if err != nil {
		fmt.Print(err.Error())
	}

	cities := []City{}
	json.Unmarshal(raw, &cities)

	newObj := new(citiesAPI)
	newObj.cities = cities

	return newObj
}
