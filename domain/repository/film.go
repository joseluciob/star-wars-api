package repository

import "star-wars-api/domain/entity"

type FilmRepository interface {
	Save(*entity.Film) (*entity.Film, map[string]string)
	Get(uint64) (*entity.Film, error)
}
