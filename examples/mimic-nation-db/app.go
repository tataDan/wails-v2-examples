package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
	var connectionError error
	if username != "root" || password != "123" {
		connectionError = errors.New("username should be 'root' and password should be '123'")
	}
	if connectionError != nil {
		errorMsg = fmt.Sprintf("%s\n%s", "Possibly incorrect username and/or password\n", connectionError.Error())
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
	var countries_read []Country

	fileName := "countries.txt"

	file, e := os.Open(fileName)
	if e != nil {
		fmt.Println("Error is = ", e)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var country Country
		s := strings.Split(scanner.Text(), ",")

		var areaAsFloat float64
		var err error
		// if areaAsFloat, err = strconv.ParseFloat(s[1], 64); err == nil {
		areaAsFloat, err = strconv.ParseFloat(s[1], 64)
		if err == nil {
			country.Name, country.Area, country.NationalDay = s[0], areaAsFloat, s[2]
		} else {
			log.Printf("Error with countries.txt file: %s", err)
		}
		countries_read = append(countries_read, country)
	}

	countries := []Country{}

	for _, v := range countries_read {
		if query_type == Like {
			if strings.HasPrefix(v.Name, country_name) {
				countries = append(countries, v)
			}
		} else if query_type == Equals {
			if v.Name == country_name {
				countries = append(countries, v)
				break
			}
		}
	}

	return countries
}
