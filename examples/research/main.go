package main

import (
	"database/sql"
	"embed"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

var db *sql.DB

func main() {
	// Open the database
	db, _ = sql.Open("sqlite3", "./reference.db")
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	defer db.Close()

	app := NewApp()

	err := wails.Run(&options.App{
		Title:            "Research",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
