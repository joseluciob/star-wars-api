package http

import (
	"star-wars-api/configs"
	"time"

	"github.com/valyala/fasthttp"
)

type Http struct {
	client *fasthttp.Client
}

func New(cfg *configs.Http) *Http {

	readTimeout, _ := time.ParseDuration("500ms")
	writeTimeout, _ := time.ParseDuration("500ms")
	maxIdleConnDuration, _ := time.ParseDuration("1h")
	client := &fasthttp.Client{
		ReadTimeout:         readTimeout,
		WriteTimeout:        writeTimeout,
		MaxIdleConnDuration: maxIdleConnDuration,
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}

	return &Http{
		client: client,
	}
}
