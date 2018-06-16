package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// City is structure for storing city.
type City struct {
	Name     string   `json:"name"`
	Location location `json:"location"`
}

type citiesAPI struct {
	cities []City
}

func HandleCity(c *gin.Context) {
	name := c.Param("name")

	c.String(http.StatusOK, "Hello %s", name)
}

func newCitiesAPI() *citiesAPI {
	raw, err := ioutil.ReadFile("resources/cities.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cities := []City{}
	json.Unmarshal(raw, &cities)

	// TODO: Finish with this method.
	// return

}
