package entity

import "time"

type Film struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:100;not null;" json:"title"`
	Director    string    `gorm:"size:100;not null;" json:"director"`
	ReleaseDate string    `gorm:"size:100;not null;" json:"release_date"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Planets     []*Planet `gorm:"many2many:planets_films;"`
}
