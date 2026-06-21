package models

import "time"

type Lån struct {
	ID                               int              `gorm:"primaryKey" json:"id"`
	Navn                             string           `gorm:"not null" json:"navn"`
	BeløpStart                       int              `gorm:"not null;column:belop_start" json:"beløp_start"`
	Rente                            int              `gorm:"not null" json:"rente"`
	Startdato                        time.Time        `gorm:"not null;default:CURRENT_TIMESTAMP" json:"startdato"`
	NedbetalingPlanlagtAvsluttetDato *time.Time       `json:"nedbetaling_planlagt_avsluttet_dato"`
	OwnerId                          *int             `gorm:"index" json:"eier_id"`
	Eier                             *Eier            `gorm:"foreignKey:OwnerId" json:"eier,omitempty"`
	Nedbetalinger                    []LånNedbetaling `gorm:"foreignKey:LånId" json:"nedbetalinger,omitempty"`
	CreatedAt                        time.Time        `json:"created_at"`
	UpdatedAt                        time.Time        `json:"updated_at"`
}

func (Lån) TableName() string {
	return "lan"
}
