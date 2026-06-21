package models

import "time"

type FondSparing struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	FondId    int       `gorm:"not null;index" json:"fond_id"`
	Fond      *Fond     `gorm:"foreignKey:FondId" json:"fond,omitempty"`
	UtgiftId  int       `gorm:"not null;index" json:"utgift_id"`
	Utgift    *Utgift   `gorm:"foreignKey:UtgiftId" json:"utgift,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (FondSparing) TableName() string {
	return "fond_sparing"
}

type KontoSparing struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	KontoId   int       `gorm:"not null;index" json:"konto_id"`
	Konto     *Konto    `gorm:"foreignKey:KontoId" json:"konto,omitempty"`
	UtgiftId  int       `gorm:"not null;index" json:"utgift_id"`
	Utgift    *Utgift   `gorm:"foreignKey:UtgiftId" json:"utgift,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (KontoSparing) TableName() string {
	return "konto_sparing"
}

type LånNedbetaling struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	LånId     int       `gorm:"not null;index;column:lan_id" json:"lån_id"`
	Lån       *Lån      `gorm:"foreignKey:LånId" json:"lån,omitempty"`
	UtgiftId  int       `gorm:"not null;index" json:"utgift_id"`
	Utgift    *Utgift   `gorm:"foreignKey:UtgiftId" json:"utgift,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (LånNedbetaling) TableName() string {
	return "lan_nedbetaling"
}
