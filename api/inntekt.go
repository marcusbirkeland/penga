package api

import (
	"context"
	"time"

	. "penga/db"
	. "penga/db/models"

	"gorm.io/gorm"
)

type InntektInput struct {
	Navn        string
	Beløp       int
	Beskrivelse string
	Dato        time.Time
	Fast        bool
	FastDato    *time.Time
	FastÅrlig   bool
	OwnerID     *int
}

func GetInntekter() ([]Inntekt, error) {
	ctx := context.Background()

	return gorm.G[Inntekt](DB).Preload("Eier", nil).Find(ctx)
}

func GetInntekt(id int) (Inntekt, error) {
	ctx := context.Background()

	return gorm.G[Inntekt](DB).Preload("Eier", nil).Where("id = ?", id).First(ctx)
}

func GetInntekterByEier(eierID int) ([]Inntekt, error) {
	ctx := context.Background()

	return gorm.G[Inntekt](DB).Preload("Eier", nil).Where("owner_id = ?", eierID).Find(ctx)
}

func CreateInntekt(input *InntektInput) (int, error) {
	ctx := context.Background()

	inntekt := Inntekt{
		Navn:        input.Navn,
		Beløp:       input.Beløp,
		Beskrivelse: input.Beskrivelse,
		Dato:        input.Dato,
		Fast:        input.Fast,
		FastDato:    input.FastDato,
		FastÅrlig:   input.FastÅrlig,
		OwnerId:     input.OwnerID,
	}

	err := gorm.G[Inntekt](DB).Create(ctx, &inntekt)
	return inntekt.ID, err
}

func UpdateInntekt(id int, input *InntektInput) error {
	ctx := context.Background()

	inntekt, err := gorm.G[Inntekt](DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return err
	}

	inntekt.Navn = input.Navn
	inntekt.Beløp = input.Beløp
	inntekt.Beskrivelse = input.Beskrivelse
	inntekt.Dato = input.Dato
	inntekt.Fast = input.Fast
	inntekt.FastDato = input.FastDato
	inntekt.FastÅrlig = input.FastÅrlig
	inntekt.OwnerId = input.OwnerID

	_, err = gorm.G[Inntekt](DB).Where("id = ?", id).Updates(ctx, inntekt)
	return err
}

func DeleteInntekt(id int) error {
	ctx := context.Background()

	_, err := gorm.G[Inntekt](DB).Where("id = ?", id).Delete(ctx)
	return err
}
