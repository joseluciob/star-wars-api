package service

import (
	"context"
	"star-wars-api/domain/entity"
)

type PlanetServiceInterface interface {
	Import(context.Context) error
	GetAll() ([]entity.Planet, error)
	GetByName(string) (*entity.Planet, error)
	Get(uint64) (*entity.Planet, error)
	Delete(uint64) error
}
