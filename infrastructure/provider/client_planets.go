package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

const RESOURCE_PLANETS = "planets"

type Planet struct {
	ID             uint64
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	ReleaseDate    string   `json:"release_date"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	Films          []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
	FilmsId        []uint64
}

type GetPlanetsResponse struct {
	Planets  []Planet `json:"results"`
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
}

func (e *SwApi) GetPlanets(ctx context.Context, page int) (*GetPlanetsResponse, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(fmt.Sprintf("%s/%s", e.config.SwApiBase, RESOURCE_PLANETS))
	req.Header.SetMethod(fasthttp.MethodGet)
	req.URI().QueryArgs().Add("page", strconv.Itoa(page))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := e.HttpClient.Do(req, resp)
	if err != nil {
		return nil, err
	}

	b := resp.Body()
	response := GetPlanetsResponse{}
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, err
	}

	for i, planet := range response.Planets {
		response.Planets[i].ID = extractIdFromURL(planet.URL, RESOURCE_PLANETS)
		for _, film := range planet.Films {
			filmId := extractIdFromURL(film, "films")
			response.Planets[i].FilmsId = append(response.Planets[i].FilmsId, filmId)
		}
	}

	return &response, nil
}
