package main

import (
	"embed"
	"fmt"
	"os"

	"database/sql"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed all:frontend/dist
var assets embed.FS

func initDb() {
	conn, err := sql.Open("sqlite3", "./db/sqlite.db")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	rows, _ := conn.Query("SELECT * from inntekt")
	defer rows.Close()
	for rows.Next() {
		var a int
		var b string
		_ = rows.Scan(&a, &b)
		fmt.Printf("%d, %s\n", a, b)
	}
}

func main() {

	initDb()
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "penga",
		Width:  1920,
		Height: 1080,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
