package handlers

import (
	"net/http"
	"star-wars-api/domain/service"
	error_h "star-wars-api/infrastructure/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Planet struct {
	planetService service.PlanetServiceInterface
}

func NewPlanet(planetService service.PlanetServiceInterface) *Planet {
	return &Planet{
		planetService: planetService,
	}
}

func (pl *Planet) GetAll(c *gin.Context) {
	var result interface{}
	var err error
	name, ok := c.GetQuery("name")
	if ok {
		result, err = pl.planetService.GetByName(name)
	} else {
		result, err = pl.planetService.GetAll()
	}
	if err != nil {
		error := error_h.ErrorHandler(err)
		c.JSON(error.Code, error.Message)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (pl *Planet) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "only numbers are accepted as an identifier parameter"})
		return
	}

	planet, err := pl.planetService.Get(id)
	if err != nil {
		error := error_h.ErrorHandler(err)
		c.JSON(error.Code, error.Message)
		return
	}
	c.JSON(http.StatusOK, planet)
}

func (pl *Planet) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "only numbers are accepted as an identifier parameter"})
		return
	}

	err = pl.planetService.Delete(id)
	if err != nil {
		error := error_h.ErrorHandler(err)
		c.JSON(error.Code, error.Message)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
