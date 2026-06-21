package main

import (
	"context"
	"fmt"

	"penga/api"
	"penga/db/models"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show timeeeee!", name)
}

func (a *App) GetEiere() ([]models.Eier, error) {
	return api.GetEiere()
}

func (a *App) CreateEier(navn string) (int, error) {
	return api.CreateEier(navn)
}

// --- Fond ---

func (a *App) GetAllFond() ([]models.Fond, error) {
	return api.GetAllFond()
}

func (a *App) GetFond(id int) (models.Fond, error) {
	return api.GetFond(id)
}

func (a *App) GetAllFondByEier(eierID int) ([]models.Fond, error) {
	return api.GetAllFondByEier(eierID)
}

func (a *App) CreateFond(input *api.FondInput) (int, error) {
	return api.CreateFond(input)
}

func (a *App) UpdateFond(id int, input *api.FondInput) error {
	return api.UpdateFond(id, input)
}

func (a *App) DeleteFond(id int) error {
	return api.DeleteFond(id)
}

func (a *App) GetAllFondBySparing() ([]models.Fond, error) {
	return api.GetAllFondBySparing()
}

// --- Inntekt ---

func (a *App) GetInntekter() ([]models.Inntekt, error) {
	return api.GetInntekter()
}

func (a *App) GetInntekt(id int) (models.Inntekt, error) {
	return api.GetInntekt(id)
}

func (a *App) GetInntekterByEier(eierID int) ([]models.Inntekt, error) {
	return api.GetInntekterByEier(eierID)
}

func (a *App) CreateInntekt(input *api.InntektInput) (int, error) {
	return api.CreateInntekt(input)
}

func (a *App) UpdateInntekt(id int, input *api.InntektInput) error {
	return api.UpdateInntekt(id, input)
}

func (a *App) DeleteInntekt(id int) error {
	return api.DeleteInntekt(id)
}

// --- Konto ---

func (a *App) GetKontoer() ([]models.Konto, error) {
	return api.GetKontoer()
}

func (a *App) GetKonto(id int) (models.Konto, error) {
	return api.GetKonto(id)
}

func (a *App) GetKontoerByEier(eierID int) ([]models.Konto, error) {
	return api.GetKontoerByEier(eierID)
}

func (a *App) CreateKonto(input *api.KontoInput) (int, error) {
	return api.CreateKonto(input)
}

func (a *App) UpdateKonto(id int, input *api.KontoInput) error {
	return api.UpdateKonto(id, input)
}

func (a *App) DeleteKonto(id int) error {
	return api.DeleteKonto(id)
}

// --- Lån ---

func (a *App) GetLån() ([]models.Lån, error) {
	return api.GetLån()
}

func (a *App) GetLånByID(id int) (models.Lån, error) {
	return api.GetLånByID(id)
}

func (a *App) GetLånByEier(eierID int) ([]models.Lån, error) {
	return api.GetLånByEier(eierID)
}

func (a *App) CreateLån(input *api.LånInput) (int, error) {
	return api.CreateLån(input)
}

func (a *App) UpdateLån(id int, input *api.LånInput) error {
	return api.UpdateLån(id, input)
}

func (a *App) DeleteLån(id int) error {
	return api.DeleteLån(id)
}

// --- Utgift ---

func (a *App) GetUtgifter() ([]models.Utgift, error) {
	return api.GetUtgifter()
}

func (a *App) GetUtgift(id int) (models.Utgift, error) {
	return api.GetUtgift(id)
}

func (a *App) GetUtgifterByEier(eierID int) ([]models.Utgift, error) {
	return api.GetUtgifterByEier(eierID)
}

func (a *App) CreateUtgift(input *api.UtgiftInput) (int, error) {
	return api.CreateUtgift(input)
}

func (a *App) UpdateUtgift(id int, input *api.UtgiftInput) error {
	return api.UpdateUtgift(id, input)
}

func (a *App) DeleteUtgift(id int) error {
	return api.DeleteUtgift(id)
}
