package models

import "time"

type Fond struct {
	ID                              int           `gorm:"primaryKey" json:"id"`
	Navn                            string        `gorm:"not null" json:"navn"`
	BeløpStart                      int           `gorm:"not null;column:belop_start" json:"beløp_start"`
	Startdato                       time.Time     `gorm:"not null" json:"startdato"`
	ForventetÅrligAvkastningProsent *int          `json:"forventet_årlig_avkastning_prosent"`
	PrisProsent                     *int          `json:"pris_prosent"`
	ISIN                            *string       `json:"ISIN"`
	OwnerId                         *int          `gorm:"index" json:"eier_id"`
	Eier                            *Eier         `gorm:"foreignKey:OwnerId" json:"eier,omitempty"`
	Sparing                         []FondSparing `gorm:"foreignKey:FondId" json:"sparing,omitempty"`
	CreatedAt                       time.Time     `json:"created_at"`
	UpdatedAt                       time.Time     `json:"updated_at"`
}

func (Fond) TableName() string {
	return "fond"
}
