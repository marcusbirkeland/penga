package models

import "time"

type Konto struct {
	ID         int            `gorm:"primaryKey" json:"id"`
	Navn       string         `gorm:"not null" json:"navn"`
	BeløpStart int            `gorm:"not null;column:belop_start" json:"beløp_start"`
	Startdato  time.Time      `gorm:"not null" json:"startdato"`
	Rente      *int           `json:"rente"`
	OwnerId    *int           `gorm:"index" json:"eier_id"`
	Eier       *Eier          `gorm:"foreignKey:OwnerId" json:"eier,omitempty"`
	Sparing    []KontoSparing `gorm:"foreignKey:KontoId" json:"sparing,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

func (Konto) TableName() string {
	return "konto"
}
