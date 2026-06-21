package api

import (
	"context"
	. "penga/db"
	. "penga/db/models"

	"gorm.io/gorm"
)

func GetEiere() ([]Eier, error) {
	ctx := context.Background()

	return gorm.G[Eier](DB).Find(ctx)
}

func CreateEier(navn string) (eier_id int, error error) {
	eier := Eier{Navn: navn}

	ctx := context.Background()
	err := gorm.G[Eier](DB).Create(ctx, &eier)
	eier_id = eier.ID
	error = err

	return
}
