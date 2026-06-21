package api

import (
	"context"
	"time"

	. "penga/db"
	. "penga/db/models"

	"gorm.io/gorm"
)

func GetUtgifter() ([]Utgift, error) {
	ctx := context.Background()

	return gorm.G[Utgift](DB).Preload("Eier", nil).Preload("FondSparing", nil).Preload("KontoSparing", nil).Preload("LånNedbetaling", nil).Find(ctx)
}

func GetUtgift(id int) (Utgift, error) {
	ctx := context.Background()

	return gorm.G[Utgift](DB).Preload("Eier", nil).Preload("FondSparing", nil).Preload("KontoSparing", nil).Preload("LånNedbetaling", nil).Where("id = ?", id).First(ctx)
}

func GetUtgifterByEier(eierID int) ([]Utgift, error) {
	ctx := context.Background()

	return gorm.G[Utgift](DB).Preload("Eier", nil).Preload("FondSparing", nil).Preload("KontoSparing", nil).Preload("LånNedbetaling", nil).Where("owner_id = ?", eierID).Find(ctx)
}

func CreateUtgift(navn string, belop int, dato time.Time, type_ *string, beskrivelse string, fast bool, fastDato *time.Time, fastArlig bool, ownerID *int) (int, error) {
	ctx := context.Background()

	utgift := Utgift{
		Navn:        navn,
		Beløp:       belop,
		Dato:        dato,
		Type:        type_,
		Beskrivelse: beskrivelse,
		Fast:        fast,
		FastDato:    fastDato,
		FastÅrlig:   fastArlig,
		OwnerId:     ownerID,
	}

	err := gorm.G[Utgift](DB).Create(ctx, &utgift)
	return utgift.ID, err
}

func UpdateUtgift(id int, navn string, belop int, dato time.Time, type_ *string, beskrivelse string, fast bool, fastDato *time.Time, fastArlig bool, ownerID *int) error {
	ctx := context.Background()

	utgift, err := gorm.G[Utgift](DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return err
	}

	utgift.Navn = navn
	utgift.Beløp = belop
	utgift.Dato = dato
	utgift.Type = type_
	utgift.Beskrivelse = beskrivelse
	utgift.Fast = fast
	utgift.FastDato = fastDato
	utgift.FastÅrlig = fastArlig
	utgift.OwnerId = ownerID

	_, err = gorm.G[Utgift](DB).Where("id = ?", id).Updates(ctx, utgift)
	return err
}

func DeleteUtgift(id int) error {
	ctx := context.Background()

	_, err := gorm.G[Utgift](DB).Where("id = ?", id).Delete(ctx)
	return err
}
