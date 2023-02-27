package provider

import "context"

type GetFilmsResponse struct {
	Title       string
	Director    string
	ReleaseDate string
}

func (e *SwApi) GetFilms(ctx context.Context, link string) *GetFilmsResponse {
}
