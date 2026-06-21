package api

import (
	"context"
	"time"

	. "penga/db"
	. "penga/db/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

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

func CreateFond(navn string, belopStart int, startdato time.Time, forventetArligAvkastningProsent *int, prisProsent *int, isin *string, ownerID *int) (int, error) {
	ctx := context.Background()

	fond := Fond{
		Navn:                            navn,
		BeløpStart:                      belopStart,
		Startdato:                       startdato,
		ForventetÅrligAvkastningProsent: forventetArligAvkastningProsent,
		PrisProsent:                     prisProsent,
		ISIN:                            isin,
		OwnerId:                         ownerID,
	}

	err := gorm.G[Fond](DB).Create(ctx, &fond)
	return fond.ID, err
}

func UpdateFond(id int, navn string, belopStart int, startdato time.Time, forventetArligAvkastningProsent *int, prisProsent *int, isin *string, ownerID *int) error {
	ctx := context.Background()

	fond, err := gorm.G[Fond](DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return err
	}

	fond.Navn = navn
	fond.BeløpStart = belopStart
	fond.Startdato = startdato
	fond.ForventetÅrligAvkastningProsent = forventetArligAvkastningProsent
	fond.PrisProsent = prisProsent
	fond.ISIN = isin
	fond.OwnerId = ownerID

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
