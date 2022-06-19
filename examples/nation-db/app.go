package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Country struct {
	Name        string
	Area        float64
	NationalDay string
}

type QueryType string

const (
	Like   QueryType = "Like"
	Equals QueryType = "Equals"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Connect(username string, password string) bool {
	var errorMsg string
	connectionError := connect(username, password)
	if connectionError != nil {
		errorMsg = fmt.Sprintf("%s\n%s", connectionError.Error(), "Possibly incorrect username and/or password")
		_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:         "ERROR CONNECTING TO DATABASE",
			Message:       errorMsg,
			Buttons:       []string{"OK"},
			DefaultButton: "OK",
		})

		if err != nil {
			log.Println(err)
			return false
		}
		return false
	}

	return true
}

func (a *App) Query(query_type QueryType, country_name string) []Country {
	var whereValue string

	db := getDb()
	if db == nil {
		log.Println("db is nil")
		return nil
	}

	if query_type == Like {
		whereValue = fmt.Sprintf("LIKE '%s%%'", country_name)
	} else if query_type == Equals {
		whereValue = fmt.Sprintf(" = '%s'", country_name)
	}
	qryStr := fmt.Sprintf(
		"SELECT name, area, IFNULL(national_day, '') FROM countries WHERE name %s", whereValue)
	rows, err := db.Query(qryStr)
	if err != nil {
		errStr := fmt.Sprintf("%s\n%s", err, "Possibly incorrect password was used.")
		log.Fatalln(errStr)
		return nil
	}
	defer rows.Close()

	countries := []Country{}

	for rows.Next() {
		country := Country{}
		if err := rows.Scan(&country.Name, &country.Area, &country.NationalDay); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		countries = append(countries, country)

		err = rows.Err()
		if err != nil {
			errStr := fmt.Sprintf("%s", err)
			log.Fatalln(errStr)
			return nil
		}
	}

	return countries
}
