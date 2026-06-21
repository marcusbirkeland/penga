package api

import (
	"context"
	"time"

	. "penga/db"
	. "penga/db/models"

	"gorm.io/gorm"
)

type LånInput struct {
	Navn                             string
	BeløpStart                       int
	Rente                            int
	Startdato                        time.Time
	NedbetalingPlanlagtAvsluttetDato *time.Time
	OwnerID                          *int
}

func GetLån() ([]Lån, error) {
	ctx := context.Background()

	return gorm.G[Lån](DB).Preload("Eier", nil).Preload("Nedbetalinger", nil).Find(ctx)
}

func GetLånByID(id int) (Lån, error) {
	ctx := context.Background()

	return gorm.G[Lån](DB).Preload("Eier", nil).Preload("Nedbetalinger", nil).Where("id = ?", id).First(ctx)
}

func GetLånByEier(eierID int) ([]Lån, error) {
	ctx := context.Background()

	return gorm.G[Lån](DB).Preload("Eier", nil).Preload("Nedbetalinger", nil).Where("owner_id = ?", eierID).Find(ctx)
}

func CreateLån(input *LånInput) (int, error) {
	ctx := context.Background()

	lan := Lån{
		Navn:                             input.Navn,
		BeløpStart:                       input.BeløpStart,
		Rente:                            input.Rente,
		Startdato:                        input.Startdato,
		NedbetalingPlanlagtAvsluttetDato: input.NedbetalingPlanlagtAvsluttetDato,
		OwnerId:                          input.OwnerID,
	}

	err := gorm.G[Lån](DB).Create(ctx, &lan)
	return lan.ID, err
}

func UpdateLån(id int, input *LånInput) error {
	ctx := context.Background()

	lan, err := gorm.G[Lån](DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return err
	}

	lan.Navn = input.Navn
	lan.BeløpStart = input.BeløpStart
	lan.Rente = input.Rente
	lan.Startdato = input.Startdato
	lan.NedbetalingPlanlagtAvsluttetDato = input.NedbetalingPlanlagtAvsluttetDato
	lan.OwnerId = input.OwnerID

	_, err = gorm.G[Lån](DB).Where("id = ?", id).Updates(ctx, lan)
	return err
}

func DeleteLån(id int) error {
	ctx := context.Background()

	_, err := gorm.G[Lån](DB).Where("id = ?", id).Delete(ctx)
	return err
}
