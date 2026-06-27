package models

import "time"

type Utgift struct {
	ID             int              `gorm:"primaryKey" json:"id"`
	Navn           string           `gorm:"not null" json:"navn"`
	Beløp          int              `gorm:"not null;column:belop" json:"beløp"`
	Dato           time.Time        `gorm:"not null" json:"dato"`
	Type           *string          `json:"type"`
	Beskrivelse    string           `json:"beskrivelse"`
	Fast           bool             `gorm:"not null" json:"fast"`
	FastDato       *time.Time       `json:"fast_dato"`
	FastÅrlig      bool             `gorm:"not null;default:false;column:fast_arlig" json:"fast_årlig"`
	OwnerId        *int             `gorm:"index" json:"eier_id"`
	Eier           *Eier            `gorm:"foreignKey:OwnerId" json:"eier,omitempty"`
	FondSparing    []FondSparing    `gorm:"foreignKey:UtgiftId" json:"fond_sparing,omitempty"`
	KontoSparing   []KontoSparing   `gorm:"foreignKey:UtgiftId" json:"konto_sparing,omitempty"`
	LånNedbetaling []LånNedbetaling `gorm:"foreignKey:UtgiftId" json:"lån_nedbetaling,omitempty"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

func (Utgift) TableName() string {
	return "utgift"
}
