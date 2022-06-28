package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Response struct {
	Message string
	Status  string
}

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) FetchAnyBreed() map[string]string {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]string
	json.Unmarshal(responseData, &data)

	return data
}

func (a *App) SelectBreed() []string {
	var breeds []string

	response, err := http.Get("https://dog.ceo/api/breeds/list/all")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]map[string][]string
	json.Unmarshal(responseData, &data)

	for _, v := range data {
		for k := range v {
			breeds = append(breeds, k)
		}
	}

	sort.Strings(breeds)

	return breeds
}

func (a *App) FetchByBreed(breed string) []string {
	var photoUrls []string

	url := fmt.Sprintf("%s%s%s%s", "https://dog.ceo/api/", "breed/", breed, "/images")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string][]string
	json.Unmarshal(responseData, &data)

	for _, v := range data {
		photoUrls = append(photoUrls, v...)
	}

	return photoUrls
}
