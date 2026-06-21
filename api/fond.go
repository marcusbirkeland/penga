package api

import (
	"context"
	"time"

	. "penga/db"
	. "penga/db/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FondInput struct {
	Navn                            string
	BeløpStart                      int
	Startdato                       time.Time
	ForventetÅrligAvkastningProsent *int
	PrisProsent                     *int
	ISIN                            *string
	OwnerID                         *int
}

func GetAllFond() ([]Fond, error) {
	ctx := context.Background()

	return gorm.G[Fond](DB).Preload("Eier", nil).Preload("Sparing", nil).Find(ctx)
}

func GetFond(id int) (Fond, error) {
	ctx := context.Background()

	return gorm.G[Fond](DB).Preload("Eier", nil).Preload("Sparing", nil).Where("id = ?", id).First(ctx)
}

func GetAllFondByEier(eierID int) ([]Fond, error) {
	ctx := context.Background()

	return gorm.G[Fond](DB).Preload("Eier", nil).Preload("Sparing", nil).Where("owner_id = ?", eierID).Find(ctx)
}

func CreateFond(input *FondInput) (int, error) {
	ctx := context.Background()

	fond := Fond{
		Navn:                            input.Navn,
		BeløpStart:                      input.BeløpStart,
		Startdato:                       input.Startdato,
		ForventetÅrligAvkastningProsent: input.ForventetÅrligAvkastningProsent,
		PrisProsent:                     input.PrisProsent,
		ISIN:                            input.ISIN,
		OwnerId:                         input.OwnerID,
	}

	err := gorm.G[Fond](DB).Create(ctx, &fond)
	return fond.ID, err
}

func UpdateFond(id int, input *FondInput) error {
	ctx := context.Background()

	fond, err := gorm.G[Fond](DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return err
	}

	fond.Navn = input.Navn
	fond.BeløpStart = input.BeløpStart
	fond.Startdato = input.Startdato
	fond.ForventetÅrligAvkastningProsent = input.ForventetÅrligAvkastningProsent
	fond.PrisProsent = input.PrisProsent
	fond.ISIN = input.ISIN
	fond.OwnerId = input.OwnerID

	_, err = gorm.G[Fond](DB).Where("id = ?", id).Updates(ctx, fond)
	return err
}

func DeleteFond(id int) error {
	ctx := context.Background()

	_, err := gorm.G[Fond](DB).Where("id = ?", id).Delete(ctx)
	return err
}

func GetAllFondBySparing() ([]Fond, error) {
	ctx := context.Background()

	return gorm.G[Fond](DB).Joins(clause.Has("Sparing"), func(db gorm.JoinBuilder, joinTable clause.Table, curTable clause.Table) error {
		return nil
	}).Preload("Eier", nil).Preload("Sparing", nil).Find(ctx)
}
