package persistence

import (
	"fmt"
	"star-wars-api/configs"
	"star-wars-api/domain/entity"
	"star-wars-api/domain/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	Planet repository.PlanetRepository
	Film   repository.FilmRepository
	db     *gorm.DB
}

func NewRepositories(dbConfig *configs.DB) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Name, dbConfig.Pass)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Planet: NewPlanetRepository(db),
		Film:   NewFilmRepository(db),
		db:     db,
	}, nil
}

func (s *Repositories) Close() {
	defer func() {
		db, err := s.db.DB()
		if err != nil {
			panic(err)
		}
		_ = db.Close()
	}()
}

func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.Film{}, &entity.Planet{})
}
