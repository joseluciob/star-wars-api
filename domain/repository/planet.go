package repository

import "star-wars-api/domain/entity"

type PlanetRepository interface {
	Save(*entity.Planet) (*entity.Planet, error)
	Update(*entity.Planet) (*entity.Planet, error)
	Get(uint64) (*entity.Planet, error)
	GetByName(string) (*entity.Planet, error)
	GetAll() ([]entity.Planet, error)
	SyncFilms(*entity.Planet, []entity.Film)
	Delete(uint64) error
}
