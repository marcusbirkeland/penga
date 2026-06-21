package models

import "time"

type Eier struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Navn      string    `gorm:"not null" json:"navn"`
	Inntekter []Inntekt `gorm:"foreignKey:OwnerId" json:"inntekter,omitempty"`
	Utgifter  []Utgift  `gorm:"foreignKey:OwnerId" json:"utgifter,omitempty"`
	Lån       []Lån     `gorm:"foreignKey:OwnerId" json:"lån,omitempty"`
	Fond      []Fond    `gorm:"foreignKey:OwnerId" json:"fond,omitempty"`
	Kontoer   []Konto   `gorm:"foreignKey:OwnerId" json:"kontoer,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Eier) TableName() string {
	return "eier"
}
