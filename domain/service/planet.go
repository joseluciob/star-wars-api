package service

import (
	"context"
	"star-wars-api/domain/entity"
	"star-wars-api/infrastructure/persistence"
	"star-wars-api/infrastructure/provider"
)

type PlanetService struct {
	repos    *persistence.Repositories
	provider provider.Provider
}

func NewPlanetService(repos *persistence.Repositories, provider provider.Provider) *PlanetService {
	return &PlanetService{repos, provider}
}

func (s *PlanetService) GetAll() ([]entity.Planet, error) {
	return s.repos.Planet.GetAll()
}

func (s *PlanetService) Get(id uint64) (*entity.Planet, error) {
	return s.repos.Planet.Get(id)
}

func (s *PlanetService) GetByName(name string) (*entity.Planet, error) {
	return s.repos.Planet.GetByName(name)
}

func (s *PlanetService) Delete(id uint64) error {
	return s.repos.Planet.Delete(id)
}

func (s *PlanetService) Import(context context.Context) {
	page := 1
	for {
		result, _ := s.provider.GetPlanets(context, page)
		for _, plan := range result.Planets {
			planet := &entity.Planet{}
			planet.ID = plan.ID
			planet.Name = plan.Name
			planet.Climate = plan.Climate
			planet.Terrain = plan.Terrain
			s.repos.Planet.Save(planet)

			var films []entity.Film
			for _, filmId := range plan.FilmsId {
				film := entity.Film{}
				film.ID = filmId
				films = append(films, film)
			}

			s.repos.Planet.SyncFilms(planet, films)
		}

		if result.Next == "" {
			break
		}
		page = page + 1
	}
}
