package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

const RESOURCE_FILMS = "films"

type Film struct {
	ID            uint64
	Title         string   `json:"title"`
	EpisodeID     int      `json:"episode_id"`
	OpeningCrawl  string   `json:"opening_crawl"`
	Director      string   `json:"director"`
	Producer      string   `json:"producer"`
	CharacterURLs []string `json:"characters"`
	PlanetURLs    []string `json:"planets"`
	StarshipURLs  []string `json:"starships"`
	VehicleURLs   []string `json:"vehicles"`
	SpeciesURLs   []string `json:"species"`
	Created       string   `json:"created"`
	Edited        string   `json:"edited"`
	URL           string   `json:"url"`
	ReleaseDate   string   `json:"release_date"`
}

type GetFilmsResponse struct {
	Films    []Film `json:"results"`
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func (e *SwApi) GetFilms(ctx context.Context, page int) (*GetFilmsResponse, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(fmt.Sprintf("%s/%s", e.config.SwApiBase, RESOURCE_FILMS))
	req.Header.SetMethod(fasthttp.MethodGet)
	req.URI().QueryArgs().Add("page", strconv.Itoa(page))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := e.HttpClient.Do(req, resp)
	if err != nil {
		return nil, err
	}

	b := resp.Body()
	response := GetFilmsResponse{}
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, err
	}

	for i, film := range response.Films {
		response.Films[i].ID = extractIdFromURL(film.URL, RESOURCE_FILMS)
	}

	return &response, nil
}
