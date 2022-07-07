package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"golang.org/x/text/message"
)

type CountryNames struct {
	Name struct {
		Common   string
		Official string
	}
}

type Country struct {
	Name struct {
		Common   string
		Official string
	}
	Cca2        string
	Cca3        string
	Independent bool
	UnMember    bool
	Currencies  map[string]map[string]string
	Capital     []string
	Region      string
	Languages   map[string]string
	Latlng      []float64
	Landlocked  bool
	Area        float64
	Flag        string
	Population  float64
	Timezones   []string
	Continents  []string
	Flags       map[string]string
	// Remaining fields are not part of the API
	LanguageList   []string
	CurrencyList   []string
	FlagList       []string
	LatlngText     []string
	AreaText       string
	PopulationText string
}

type Population struct {
	Name struct {
		Common   string
		Official string
	}
	Population float64
}

type Area struct {
	Name struct {
		Common   string
		Official string
	}
	Area float64
}

type PopAndArea struct {
	Name struct {
		Common   string
		Official string
	}
	Population float64
	Area       float64
}

type Density struct {
	Name    string
	Density float64
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

func (a *App) GetCountryNames() []string {
	var commonNames []string
	url := "https://restcountries.com/v3.1/all"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var countryNames []CountryNames
	decodeErr := json.Unmarshal(responseData, &countryNames)
	if decodeErr != nil {
		log.Fatal(err)
	}

	for _, name := range countryNames {
		commonNames = append(commonNames, name.Name.Common)
	}

	sort.Strings(commonNames)

	return commonNames
}

func (a *App) GetCountry(countryName string) Country {
	url := fmt.Sprintf("%s%s%s", "https://restcountries.com/v3.1/name/", countryName, "?fullText=true")

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var countries []Country
	decodeErr := json.Unmarshal(responseData, &countries)
	if decodeErr != nil {
		log.Fatal(err)
	}

	if len(countries) > 0 {
		country := countries[0]

		for _, v := range country.Languages {
			v = " " + v
			country.LanguageList = append(country.LanguageList, v)
		}

		for _, v := range country.Currencies {
			for _, v2 := range v {
				v2 = " " + v2
				country.CurrencyList = append(country.CurrencyList, v2)
			}
		}

		for _, v := range country.Flags {
			v = " " + v
			country.FlagList = append(country.FlagList, v)
		}

		for i, v := range country.Latlng {
			var latText string
			var lngText string
			if i == 0 {
				latText = fmt.Sprintf("%s: %f", "Latitude", v)
			} else if i == 1 {
				lngText = fmt.Sprintf("%s: %f", "Longitude", v)
			}
			text := fmt.Sprintf("%s %s", latText, lngText)
			country.LatlngText = append(country.LatlngText, text)
		}

		p := message.NewPrinter(message.MatchLanguage("en"))
		country.AreaText = p.Sprintf("%.0f square kilometers", country.Area)
		country.PopulationText = p.Sprintf("%.0f", country.Population)

		return country
	}
	return Country{}
}

func (a *App) GetPopulations(highToLow bool) []string {
	var populationsSorted []string
	url := "https://restcountries.com/v3.1/all"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var populations []Population
	decodeErr := json.Unmarshal(responseData, &populations)
	if decodeErr != nil {
		log.Println(err)
		// log.Fatal(err)
	}

	sort.SliceStable(populations, func(i, j int) bool {
		if highToLow {
			return populations[i].Population > populations[j].Population
		} else {
			return populations[i].Population < populations[j].Population
		}
	})

	p := message.NewPrinter(message.MatchLanguage("en"))

	for _, v := range populations {
		popLine := p.Sprintf("%s: %.0f", v.Name.Common, v.Population)
		populationsSorted = append(populationsSorted, popLine)
	}

	return populationsSorted
}

func (a *App) GetAreas(highToLow bool) []string {
	var areasSorted []string
	url := "https://restcountries.com/v3.1/all"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var areas []Area
	decodeErr := json.Unmarshal(responseData, &areas)
	if decodeErr != nil {
		log.Println(err)
		// log.Fatal(err)
	}

	sort.SliceStable(areas, func(i, j int) bool {
		if highToLow {
			return areas[i].Area > areas[j].Area
		} else {
			return areas[i].Area < areas[j].Area
		}
	})

	p := message.NewPrinter(message.MatchLanguage("en"))

	for _, v := range areas {
		areaLine := p.Sprintf("%s: %.0f", v.Name.Common, v.Area)
		areasSorted = append(areasSorted, areaLine)
	}

	return areasSorted
}

func (a *App) GetDensities(highToLow bool) []string {
	var densitiesSorted []string
	url := "https://restcountries.com/v3.1/all"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var popAndAreas []PopAndArea
	decodeErr := json.Unmarshal(responseData, &popAndAreas)
	if decodeErr != nil {
		log.Println(err)
		// log.Fatal(err)
	}

	var densities []Density
	for _, v := range popAndAreas {
		temp := Density{
			Name:    v.Name.Common,
			Density: (v.Population / v.Area),
		}
		densities = append(densities, temp)
	}

	sort.SliceStable(densities, func(i, j int) bool {
		if highToLow {
			return densities[i].Density > densities[j].Density
		} else {
			return densities[i].Density < densities[j].Density
		}
	})

	p := message.NewPrinter(message.MatchLanguage("en"))

	for _, v := range densities {
		denLine := p.Sprintf("%s: %f", v.Name, v.Density)
		densitiesSorted = append(densitiesSorted, denLine)
	}

	return densitiesSorted
}
