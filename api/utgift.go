package api

import (
	"context"
	"time"

	. "penga/db"
	. "penga/db/models"

	"gorm.io/gorm"
)

type UtgiftInput struct {
	Navn        string
	Beløp       int
	Dato        time.Time
	Type        *string
	Beskrivelse string
	Fast        bool
	FastDato    *time.Time
	FastÅrlig   bool
	OwnerID     *int
}

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

func CreateUtgift(input *UtgiftInput) (int, error) {
	ctx := context.Background()

	utgift := Utgift{
		Navn:        input.Navn,
		Beløp:       input.Beløp,
		Dato:        input.Dato,
		Type:        input.Type,
		Beskrivelse: input.Beskrivelse,
		Fast:        input.Fast,
		FastDato:    input.FastDato,
		FastÅrlig:   input.FastÅrlig,
		OwnerId:     input.OwnerID,
	}

	err := gorm.G[Utgift](DB).Create(ctx, &utgift)
	return utgift.ID, err
}

func UpdateUtgift(id int, input *UtgiftInput) error {
	ctx := context.Background()

	utgift, err := gorm.G[Utgift](DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return err
	}

	utgift.Navn = input.Navn
	utgift.Beløp = input.Beløp
	utgift.Dato = input.Dato
	utgift.Type = input.Type
	utgift.Beskrivelse = input.Beskrivelse
	utgift.Fast = input.Fast
	utgift.FastDato = input.FastDato
	utgift.FastÅrlig = input.FastÅrlig
	utgift.OwnerId = input.OwnerID

	_, err = gorm.G[Utgift](DB).Where("id = ?", id).Updates(ctx, utgift)
	return err
}

func DeleteUtgift(id int) error {
	ctx := context.Background()

	_, err := gorm.G[Utgift](DB).Where("id = ?", id).Delete(ctx)
	return err
}
