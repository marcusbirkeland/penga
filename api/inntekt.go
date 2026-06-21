package api

import (
	"context"
	"time"

	. "penga/db"
	. "penga/db/models"

	"gorm.io/gorm"
)

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

func CreateInntekt(navn string, belop int, beskrivelse string, dato time.Time, fast bool, fastDato *time.Time, fastArlig bool, ownerID *int) (int, error) {
	ctx := context.Background()

	inntekt := Inntekt{
		Navn:        navn,
		Beløp:       belop,
		Beskrivelse: beskrivelse,
		Dato:        dato,
		Fast:        fast,
		FastDato:    fastDato,
		FastÅrlig:   fastArlig,
		OwnerId:     ownerID,
	}

	err := gorm.G[Inntekt](DB).Create(ctx, &inntekt)
	return inntekt.ID, err
}

func UpdateInntekt(id int, navn string, belop int, beskrivelse string, dato time.Time, fast bool, fastDato *time.Time, fastArlig bool, ownerID *int) error {
	ctx := context.Background()

	inntekt, err := gorm.G[Inntekt](DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return err
	}

	inntekt.Navn = navn
	inntekt.Beløp = belop
	inntekt.Beskrivelse = beskrivelse
	inntekt.Dato = dato
	inntekt.Fast = fast
	inntekt.FastDato = fastDato
	inntekt.FastÅrlig = fastArlig
	inntekt.OwnerId = ownerID

	_, err = gorm.G[Inntekt](DB).Where("id = ?", id).Updates(ctx, inntekt)
	return err
}

func DeleteInntekt(id int) error {
	ctx := context.Background()

	_, err := gorm.G[Inntekt](DB).Where("id = ?", id).Delete(ctx)
	return err
}
