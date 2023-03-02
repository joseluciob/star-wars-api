package persistence

import (
	"star-wars-api/domain/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FilmRepo struct {
	db *gorm.DB
}

func NewFilmRepository(db *gorm.DB) *FilmRepo {
	return &FilmRepo{db}
}

func (r *FilmRepo) Get(id uint64) (*entity.Film, error) {
	var film entity.Film
	err := r.db.Where("id = ?", id).Take(&film).Error
	if err != nil {
		return nil, err
	}

	return &film, nil
}

func (r *FilmRepo) Save(film *entity.Film) (*entity.Film, error) {

	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"title", "director", "release_date"}),
	}).Create(&film).Error
	if err != nil {
		return nil, err
	}

	return film, nil
}

func (r *FilmRepo) GetAll() ([]entity.Film, error) {

	return nil, nil
}

func (r *FilmRepo) Update(film *entity.Film) (*entity.Film, error) {
	err := r.db.Save(&film).Error
	if err != nil {
		return nil, err
	}
	return film, nil
}
