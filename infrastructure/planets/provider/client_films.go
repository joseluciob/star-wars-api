package provider

import (
	"context"
)

type GetFilmsResponse struct {
	Title       string `json:"title"`
	Director    string `json:"director"`
	ReleaseDate string `json:"release_date"`
	URL         string `json:"url"`
}

func (e *SwApi) GetFilms(ctx context.Context, link string) *GetFilmsResponse {
	return &GetFilmsResponse{}
}
