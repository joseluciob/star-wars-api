package persistence

import (
	"star-wars-api/domain/entity"

	"gorm.io/gorm"
)

type FilmRepo struct {
	db *gorm.DB
}

func NewFilmRepository(db *gorm.DB) *FilmRepo {
	return &FilmRepo{db}
}

func (r *FilmRepo) Get(id uint64) (*entity.Film, error) {
	return nil, nil
}

func (r *FilmRepo) Save(film *entity.Film) (*entity.Film, map[string]string) {

	return nil, nil
}

func (r *FilmRepo) GetAll() ([]entity.Film, error) {

	return nil, nil
}
