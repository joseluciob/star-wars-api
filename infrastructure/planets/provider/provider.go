package provider

import (
	"context"
	"star-wars-api/configs"

	"github.com/valyala/fasthttp"
)

type Provider interface {
	GetPlanets(ctx context.Context, page int) (*GetPlanetsResponse, error)
	GetFilms(ctx context.Context, film string) (*GetFilmsResponse, error)
}
type SwApi struct {
	config     *configs.Configs
	HttpClient *fasthttp.Client
}

func New(cfg *configs.Configs) (*SwApi, error) {
	httpConfig := cfg.Http
	return &SwApi{
		config: cfg,
		HttpClient: &fasthttp.Client{
			ReadTimeout:         httpConfig.ReadTimeout,
			WriteTimeout:        httpConfig.WriteTimeout,
			MaxIdleConnDuration: httpConfig.MaxIdleConnDuration,
			Dial: (&fasthttp.TCPDialer{
				Concurrency:      httpConfig.DialConcurrency,
				DNSCacheDuration: httpConfig.DialDnsCacheDuration,
			}).Dial,
		},
	}, nil
}
