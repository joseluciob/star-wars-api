package repository

import "star-wars-api/domain/entity"

type PlanetRepository interface {
	Save(*entity.Planet) (*entity.Planet, map[string]string)
	Get(uint64) (*entity.Planet, error)
	GetAll() ([]entity.Planet, error)
}
