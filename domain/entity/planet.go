package entity

import (
	"time"

	"gorm.io/gorm"
)

type Planet struct {
	ID        uint64          `gorm:"primary_key;auto_increment" json:"id"`
	Name      string          `gorm:"size:100;not null;" json:"name"`
	Climate   string          `gorm:"size:100;not null;" json:"climate"`
	Terrain   string          `gorm:"size:100;not null;" json:"terrain"`
	Films     []Film          `gorm:"many2many:planets_films" json:"films"`
	CreatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

type PlanetFilm struct {
	PlanetId int `gorm:"primaryKey"`
	FilmId   int `gorm:"primaryKey"`
}
