package repository

import "star-wars-api/domain/entity"

type FilmRepository interface {
	Save(*entity.Film) (*entity.Film, error)
	Update(*entity.Film) (*entity.Film, error)
	Get(uint64) (*entity.Film, error)
}
