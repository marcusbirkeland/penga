package api

import (
	"context"
	"time"

	. "penga/db"
	. "penga/db/models"

	"gorm.io/gorm"
)

type KontoInput struct {
	Navn       string
	BeløpStart int
	Startdato  time.Time
	Rente      *int
	OwnerID    *int
}

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

func CreateKonto(input *KontoInput) (int, error) {
	ctx := context.Background()

	konto := Konto{
		Navn:       input.Navn,
		BeløpStart: input.BeløpStart,
		Startdato:  input.Startdato,
		Rente:      input.Rente,
		OwnerId:    input.OwnerID,
	}

	err := gorm.G[Konto](DB).Create(ctx, &konto)
	return konto.ID, err
}

func UpdateKonto(id int, input *KontoInput) error {
	ctx := context.Background()

	konto, err := gorm.G[Konto](DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return err
	}

	konto.Navn = input.Navn
	konto.BeløpStart = input.BeløpStart
	konto.Startdato = input.Startdato
	konto.Rente = input.Rente
	konto.OwnerId = input.OwnerID

	_, err = gorm.G[Konto](DB).Where("id = ?", id).Updates(ctx, konto)
	return err
}

func DeleteKonto(id int) error {
	ctx := context.Background()

	_, err := gorm.G[Konto](DB).Where("id = ?", id).Delete(ctx)
	return err
}
