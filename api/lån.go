package api

import (
	"context"
	"time"

	. "penga/db"
	. "penga/db/models"

	"gorm.io/gorm"
)

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

func CreateLån(navn string, belopStart int, rente int, startdato time.Time, nedbetalingPlanlagtAvsluttetDato *time.Time, ownerID *int) (int, error) {
	ctx := context.Background()

	lan := Lån{
		Navn:                             navn,
		BeløpStart:                       belopStart,
		Rente:                            rente,
		Startdato:                        startdato,
		NedbetalingPlanlagtAvsluttetDato: nedbetalingPlanlagtAvsluttetDato,
		OwnerId:                          ownerID,
	}

	err := gorm.G[Lån](DB).Create(ctx, &lan)
	return lan.ID, err
}

func UpdateLån(id int, navn string, belopStart int, rente int, startdato time.Time, nedbetalingPlanlagtAvsluttetDato *time.Time, ownerID *int) error {
	ctx := context.Background()

	lan, err := gorm.G[Lån](DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return err
	}

	lan.Navn = navn
	lan.BeløpStart = belopStart
	lan.Rente = rente
	lan.Startdato = startdato
	lan.NedbetalingPlanlagtAvsluttetDato = nedbetalingPlanlagtAvsluttetDato
	lan.OwnerId = ownerID

	_, err = gorm.G[Lån](DB).Where("id = ?", id).Updates(ctx, lan)
	return err
}

func DeleteLån(id int) error {
	ctx := context.Background()

	_, err := gorm.G[Lån](DB).Where("id = ?", id).Delete(ctx)
	return err
}
