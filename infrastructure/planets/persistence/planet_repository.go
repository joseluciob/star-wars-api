package persistence

import (
	"star-wars-api/domain/entity"

	"gorm.io/gorm"
)

type PlanetRepo struct {
	db *gorm.DB
}

func NewPlanetRepository(db *gorm.DB) *PlanetRepo {
	return &PlanetRepo{db}
}

func (r *PlanetRepo) Get(id uint64) (*entity.Planet, error) {
	return nil, nil
}

func (r *PlanetRepo) Save(planet *entity.Planet) (*entity.Planet, map[string]string) {

	return nil, nil
}

func (r *PlanetRepo) GetAll() ([]entity.Planet, error) {

	return nil, nil
}
