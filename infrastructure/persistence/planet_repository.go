package persistence

import (
	"star-wars-api/domain/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PlanetRepo struct {
	db *gorm.DB
}

func NewPlanetRepository(db *gorm.DB) *PlanetRepo {
	return &PlanetRepo{db}
}

func (r *PlanetRepo) Get(id uint64) (*entity.Planet, error) {
	var planet entity.Planet
	err := r.db.Preload("Films").Where("id = ?", id).Take(&planet).Error
	if err != nil {
		return nil, err
	}

	return &planet, nil
}

func (r *PlanetRepo) GetByName(name string) (*entity.Planet, error) {
	var planet entity.Planet
	err := r.db.Preload("Films").Where("name = ?", name).Take(&planet).Error
	if err != nil {
		return nil, err
	}

	return &planet, nil
}

func (r *PlanetRepo) Save(planet *entity.Planet) (*entity.Planet, error) {
	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "climate", "terrain"}),
	}).Create(&planet).Error
	if err != nil {
		return nil, err
	}

	return planet, nil
}

func (r *PlanetRepo) GetAll() ([]entity.Planet, error) {
	var planets []entity.Planet
	err := r.db.Preload("Films").Order("id asc").Find(&planets).Error
	if err != nil {
		return nil, err
	}
	return planets, nil
}

func (r *PlanetRepo) Update(planet *entity.Planet) (*entity.Planet, error) {
	err := r.db.Save(&planet).Error
	if err != nil {
		return nil, err
	}
	return planet, nil
}

func (r *PlanetRepo) SyncFilms(planet *entity.Planet, films []entity.Film) {
	r.db.Model(planet).Association("Films").Replace(films)
}

func (r *PlanetRepo) Delete(id uint64) error {
	var planet entity.Planet
	err := r.db.Where("id = ?", id).Delete(&planet).Error
	if err != nil {
		return err
	}
	return nil
}
