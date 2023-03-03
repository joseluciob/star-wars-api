package handlers

import (
	"net/http"
	"star-wars-api/domain/service"
	error_h "star-wars-api/infrastructure/httputil"
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

// GetAll lists all existing planets
//
//	@Summary      List planets
//	@Description  get planets
//	@Tags         planets
//	@Accept       json
//	@Produce      json
//	@Param		  name	query	 string	false "Search planet by name"
//	@Success      200  {array}   entity.Planet
//	@Failure      400  {object}  httputil.Error
//	@Failure      404  {object}  httputil.Error
//	@Failure      500  {object}  httputil.Error
//	@Router       /planets [get]
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

// Get get a planet by identifier
//
//	@Summary		Show an planet
//	@Description	get planet by ID
//	@Tags			planets
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Planet ID"
//	@Success		200	{object}	entity.Planet
//	@Failure		400	{object}	httputil.Error
//	@Failure		404	{object}	httputil.Error
//	@Failure		500	{object}	httputil.Error
//	@Router			/planets/{id} [get]
func (pl *Planet) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	planet, err := pl.planetService.Get(id)
	if err != nil {
		error := error_h.ErrorHandler(err)
		c.JSON(error.Code, error.Message)
		return
	}
	c.JSON(http.StatusOK, planet)
}

// Delete delete planet by id
//
//	@Summary		Delete an planet
//	@Description	Delete by planet ID
//	@Tags			planets
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Planet ID"	Format(int64)
//	@Success		204	{object}	entity.Planet
//	@Failure		400	{object}	httputil.Error
//	@Failure		404	{object}	httputil.Error
//	@Failure		500	{object}	httputil.Error
//	@Router			/planets/{id} [delete]
func (pl *Planet) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := pl.planetService.Delete(id)
	if err != nil {
		error := error_h.ErrorHandler(err)
		c.JSON(error.Code, error.Message)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
