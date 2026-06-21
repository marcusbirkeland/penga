package models

import "time"

type Inntekt struct {
	ID          int        `gorm:"primaryKey" json:"id"`
	Navn        string     `gorm:"not null" json:"navn"`
	Beløp       int        `gorm:"not null;column:belop" json:"beløp"`
	Beskrivelse string     `json:"beskrivelse"`
	Dato        time.Time  `gorm:"not null" json:"dato"`
	Fast        bool       `gorm:"not null;default:true" json:"fast"`
	FastDato    *time.Time `json:"fast_dato"`
	FastÅrlig   bool       `gorm:"not null;default:false;column:fast_arlig" json:"fast_årlig"`
	OwnerId     *int       `gorm:"index" json:"eier_id"`
	Eier        *Eier      `gorm:"foreignKey:OwnerId" json:"eier,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (Inntekt) TableName() string {
	return "inntekt"
}
