package api

import (
	"context"
	"time"

	. "penga/db"
	. "penga/db/models"

	"gorm.io/gorm"
)

func GetKontoer() ([]Konto, error) {
	ctx := context.Background()

	return gorm.G[Konto](DB).Preload("Eier", nil).Preload("Sparing", nil).Find(ctx)
}

func GetKonto(id int) (Konto, error) {
	ctx := context.Background()

	return gorm.G[Konto](DB).Preload("Eier", nil).Preload("Sparing", nil).Where("id = ?", id).First(ctx)
}

func GetKontoerByEier(eierID int) ([]Konto, error) {
	ctx := context.Background()

	return gorm.G[Konto](DB).Preload("Eier", nil).Preload("Sparing", nil).Where("owner_id = ?", eierID).Find(ctx)
}

func CreateKonto(navn string, belopStart int, startdato time.Time, rente *int, ownerID *int) (int, error) {
	ctx := context.Background()

	konto := Konto{
		Navn:       navn,
		BeløpStart: belopStart,
		Startdato:  startdato,
		Rente:      rente,
		OwnerId:    ownerID,
	}

	err := gorm.G[Konto](DB).Create(ctx, &konto)
	return konto.ID, err
}

func UpdateKonto(id int, navn string, belopStart int, startdato time.Time, rente *int, ownerID *int) error {
	ctx := context.Background()

	konto, err := gorm.G[Konto](DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return err
	}

	konto.Navn = navn
	konto.BeløpStart = belopStart
	konto.Startdato = startdato
	konto.Rente = rente
	konto.OwnerId = ownerID

	_, err = gorm.G[Konto](DB).Where("id = ?", id).Updates(ctx, konto)
	return err
}

func DeleteKonto(id int) error {
	ctx := context.Background()

	_, err := gorm.G[Konto](DB).Where("id = ?", id).Delete(ctx)
	return err
}
