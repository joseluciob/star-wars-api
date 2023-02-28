package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/valyala/fasthttp"
)

type Planet struct {
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
}

type GetPlanetsResponse struct {
	Planets  []Planet `json:"results"`
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
}

func (e *SwApi) GetPlanets(ctx context.Context, page int) *GetPlanetsResponse {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(fmt.Sprintf("%s/%s", e.config.SwApiBase, "/planets"))
	req.Header.SetMethod(fasthttp.MethodGet)
	req.URI().QueryArgs().Add("page", strconv.Itoa(page))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := e.HttpClient.Do(req, resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
		// log
	}

	b := resp.Body()
	response := GetPlanetsResponse{}
	if err := json.Unmarshal(b, &response); err != nil {
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
		// log
	}

	return &response
}
