package service

import (
	"context"
	"star-wars-api/domain/entity"
	"star-wars-api/infrastructure/persistence"
	"star-wars-api/infrastructure/provider"
	"time"
)

type FilmService struct {
	repos    *persistence.Repositories
	provider provider.Provider
}

func NewFilmService(repos *persistence.Repositories, provider provider.Provider) *FilmService {
	return &FilmService{repos, provider}
}

func (s *FilmService) Import(context context.Context) {
	page := 1
	for {
		result, _ := s.provider.GetFilms(context, page)
		for _, filmResult := range result.Films {
			film := &entity.Film{}
			film.ID = filmResult.ID
			film.Director = filmResult.Director
			film.Title = filmResult.Title
			film.ReleaseDate, _ = time.Parse("2006-01-02", filmResult.ReleaseDate)
			s.repos.Film.Save(film)
		}
		if result.Next == "" {
			break
		}
		page = page + 1
	}
}
