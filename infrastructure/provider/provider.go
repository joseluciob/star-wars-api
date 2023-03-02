package provider

import (
	"context"
	"regexp"
	"star-wars-api/configs"
	"strconv"

	"github.com/valyala/fasthttp"
)

type Provider interface {
	GetPlanets(ctx context.Context, page int) (*GetPlanetsResponse, error)
	GetFilms(ctx context.Context, page int) (*GetFilmsResponse, error)
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

func extractIdFromURL(url, resource string) uint64 {
	re := regexp.MustCompile(`/(.*)` + resource + `\/(\d*)/`)
	match := re.FindStringSubmatch(url)

	if len(match) >= 2 {
		value, _ := strconv.ParseUint(match[2], 10, 64)
		return value
	}

	return 0
}
