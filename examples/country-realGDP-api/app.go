package main

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"

	"sort"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/wcharczuk/go-chart/v2"
)

type RealGDP struct {
	Data struct {
		Values []float64
		Dates  []string
		Status []string
	}
}

type Country []struct {
	Name struct {
		Common   string
		Official string
	}
	Cca2 string
}

type Rgdp struct {
	Value  float64
	Date   string
	Status string
}

type App struct {
	ctx context.Context
}

var nameToCode map[string]string = make(map[string]string)
var pathToLineChartFile string

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetCountryCode(selectedCountry string) string {
	return nameToCode[selectedCountry]
}

func (a *App) GetCountryNames() []string {

	var countryNames []string

	response, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		fmt.Print(err.Error())
		// log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var country Country
	json.Unmarshal(responseData, &country)

	for _, v := range country {
		countryNames = append(countryNames, v.Name.Common)
	}

	sort.Strings(countryNames)

	for _, v := range country {
		nameToCode[v.Name.Common] = v.Cca2
	}

	return countryNames
}

func (a *App) GetRGDP(countryCode string) []Rgdp {
	if countryCode == "GB" {
		countryCode = "UK"
	}
	url := fmt.Sprintf("https://www.econdb.com/api/series/RGDP%s/?format=json", countryCode)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		// log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var realGDP RealGDP
	json.Unmarshal(responseData, &realGDP)

	rgdps := []Rgdp{}

	var rgdp Rgdp

	if (len(realGDP.Data.Values) == len(realGDP.Data.Dates)) && (len(realGDP.Data.Dates) == len(realGDP.Data.Status)) {
		for i := 0; i < len(realGDP.Data.Values); i++ {
			rgdp.Value = realGDP.Data.Values[i]
			rgdp.Date = realGDP.Data.Dates[i]
			rgdp.Status = realGDP.Data.Status[i]
			rgdps = append(rgdps, rgdp)
		}
	} else {
		log.Println("Error found in data for a country -- lengths of arrays are not equal")
	}

	if len(rgdps) > 0 {
		runtime.EventsEmit(a.ctx, "data-found")
	}

	return rgdps
}

func (a *App) OneCountryLineChart(selectedCountry string, countryCode string) string {
	if countryCode == "GB" {
		countryCode = "UK"
	}
	url := fmt.Sprintf("https://www.econdb.com/api/series/RGDP%s/?format=json", countryCode)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		// log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var realGDP RealGDP
	json.Unmarshal(responseData, &realGDP)

	const numValuesRequested = 10

	if (len(realGDP.Data.Values) < numValuesRequested) || (len(realGDP.Data.Dates) < numValuesRequested) {
		return ""
	}

	if len(realGDP.Data.Values) != len(realGDP.Data.Dates) {
		return ""
	}

	var values []float64
	for _, v := range realGDP.Data.Values {
		values = append(values, float64(v))
	}

	var labels []string
	for _, v := range realGDP.Data.Dates {
		labels = append(labels, v)
	}

	var xValues []float64
	for i := 0; i < numValuesRequested; i++ {
		xValues = append(xValues, float64(i+1))
	}

	var xTicks []chart.Tick
	for i := 0; i < numValuesRequested; i++ {
		xTicks = append(xTicks, chart.Tick{
			Value: float64(i + 1),
			Label: labels[len(labels)-(numValuesRequested-i)],
		})
	}

	graph := chart.Chart{
		Title: selectedCountry,
		YAxis: chart.YAxis{
			Name: "Real Gross Domestic Product",
		},
		XAxis: chart.XAxis{
			Ticks: xTicks,
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: values[len(values)-numValuesRequested:],
			},
		},
	}

	workDir, errWd := os.Getwd()
	if errWd != nil {
		log.Println(errWd)
	}
	pathToLineChartFile = filepath.Join(workDir, "lineChart.png")
	f, errCreate := os.Create(pathToLineChartFile)
	defer f.Close()
	if errCreate != nil {
		log.Println(errCreate)
	}
	graph.Render(chart.PNG, f)

	return pathToLineChartFile
}

func (a *App) OpenLineChart(pathToLineChartFile string, viewer string) {
	var errorMsg string
	path := pathToLineChartFile
	cmd := exec.Command(viewer, path)
	err := cmd.Run()

	if err != nil {
		log.Println(err)
		errorMsg = fmt.Sprintf("%s\n%s", "Could not find requested application\n", err)
		_, errMD := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:         "ERROR FINDING REQUESTED APPLICATION",
			Message:       errorMsg,
			Buttons:       []string{"OK"},
			DefaultButton: "OK",
		})

		if errMD != nil {
			log.Println(err)
		}
	}
}
