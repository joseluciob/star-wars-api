package provider

import (
	"context"
	"star-wars-api/configs"
	"star-wars-api/infrastructure/common/http"
)

type Provider interface {
	GetPlanets(ctx context.Context) (*GetPlanetsResponse, error)
	GetFilms(ctx context.Context, film string) (*GetFilmsResponse, error)
}
type SwApi struct {
	config *configs.Http
	http   *http.Http
}

func New(cfg *configs.Http) (*SwApi, error) {
	return &SwApi{
		config: cfg,
		http:   http.New(cfg),
	}, nil
}
