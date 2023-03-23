package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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

type dataToShow struct {
	Id                int
	Name              string
	City              string
	State             string
	StudentSize       int
	GradStudents      int
	TuitionInState    int
	TuitionOutOfState int
	SchoolUrl         string
	PriceCalculatoUrl string
}

type Metadata struct {
	Page     int
	Total    int
	Per_page int
}

type Result struct {
	Name                     string `json:"latest.school.name"`
	City                     string `json:"latest.school.city"`
	State                    string `json:"latest.school.state"`
	SchoolUrl                string `json:"latest.school.school_url"`
	SchoolPriceCalculatorUrl string `json:"latest.school.price_calculator_url"`
	StudentSize              int    `json:"latest.student.size"`
	GradStudents             int    `json:"latest.student.grad_students"`
	TuitionInState           int    `json:"latest.cost.tuition.in_state"`
	TuitionOutOfState        int    `json:"latest.cost.tuition.out_of_state"`
}

type CollegeData struct {
	Metadata Metadata
	Results  []Result
}

var titleToCode = make(map[string]string)
var programTitles = []string{}
var stateNameToCode = make(map[string]string)
var stateNames = []string{}

func (a *App) LoadStateData() {
	records, err := readData("name-abbr.csv")
	if err != nil {
		errMsg := fmt.Sprintf("%s  Error: %s.  %s", "States were not loaded.", err, "Application will terminate!")
		_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:         "ERROR",
			Message:       errMsg,
			Buttons:       []string{"OK"},
			DefaultButton: "OK",
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Fatal(errMsg)
	}
	for _, record := range records {
		name := record[0]
		code := record[1]
		stateNameToCode[name] = code
		stateNames = append(stateNames, name)
	}
	sort.Strings(stateNames)
}

func (a *App) GetStateNameToCode() map[string]string {
	return stateNameToCode
}

func (a *App) GetStateNames() []string {
	return stateNames
}

func (a *App) GetStateCode(name string) string {
	return stateNameToCode[name]
}

func (a *App) LoadProgramData() {
	records, err := readData("program-code.csv")
	if err != nil {
		errMsg := fmt.Sprintf("%s  Error: %s.  %s", "Program titles were not loaded.", err, "Application will terminate!")
		_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:         "ERROR",
			Message:       errMsg,
			Buttons:       []string{"OK"},
			DefaultButton: "OK",
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Fatal(errMsg)
	}
	for _, record := range records {
		title := record[0]
		code := record[1]
		programTitles = append(programTitles, title)
		titleToCode[title] = code
	}

	programTitles = removeDuplicateValues(programTitles)
	sort.Strings(programTitles)
}

func (a *App) GetProgramTitles() []string {
	return programTitles
}

func (a *App) GetProgramCode(title string) string {
	return titleToCode[title]
}

func (a *App) SearchProgramTitles(str string) []string {
	matchingTitles := []string{}
	str = strings.Trim(str, " ")
	for _, pt := range programTitles {
		if strings.Contains(strings.ToUpper(pt), strings.ToUpper(str)) {
			matchingTitles = append(matchingTitles, pt)
		}
	}
	return matchingTitles
}

func (a *App) GetData(city string, stateCode string, programCode string) []dataToShow {
	var dataIn CollegeData
	queryParams := ""
	dataOut := []dataToShow{}
	page := 0

	city = strings.Trim(city, " ")
	if city == "" && stateCode == "" && programCode == "" {
		_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:         "ERROR",
			Message:       "Please enter at least one of the following fields: city, state, program.",
			Buttons:       []string{"OK"},
			DefaultButton: "OK",
		})
		if err != nil {
			log.Println(err)
			return dataOut
		}
		return dataOut
	}

	if len(city) != 0 {
		city = strings.ReplaceAll(city, " ", "%20")
		queryParams += "&latest.school.city=" + city
	}

	if len(stateCode) != 0 {
		queryParams += "&latest.school.state=" + stateCode
	}

	if len(programCode) != 0 {
		queryParams += "&latest.programs.cip_4_digit.code=" + programCode
	}

	url := "https://api.data.gov/ed/collegescorecard/v1/schools?fields=latest.school.name,latest.school.city,latest.school.state,latest.school.school_url," +
		"latest.school.price_calculator_url,latest.student.size,latest.student.grad_students,latest.cost.tuition.in_state," +
		"latest.cost.tuition.out_of_state"
	apiKey := "&api_key=XXXXXXXXXX"
	url = fmt.Sprintf("%s%s%s", url, queryParams, apiKey)

	for {
		url := fmt.Sprintf("%s&page=%d", url, page)
		response, err := http.Get(url)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		if response.StatusCode == 403 {
			errMsg := fmt.Sprintf("%s.  %s", "StatusForbidden (403), (Possibly the api key is invalid). Please report this error to the developer", "Application will terminate!")
			_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
				Title:         "ERROR",
				Message:       errMsg,
				Buttons:       []string{"OK"},
				DefaultButton: "OK",
			})
			if err != nil {
				log.Println(err)
				return dataOut
			}
			log.Fatal(errMsg)
		}

		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(responseData, &dataIn)

		if dataIn.Metadata.Total == 0 {
			return dataOut
		}

		if len(dataIn.Results) == 0 {
			break
		}

		for _, r := range dataIn.Results {
			school := dataToShow{
				Name:              r.Name,
				City:              r.City,
				State:             r.State,
				StudentSize:       r.StudentSize,
				GradStudents:      r.GradStudents,
				SchoolUrl:         r.SchoolUrl,
				PriceCalculatoUrl: r.SchoolPriceCalculatorUrl,
				TuitionInState:    r.TuitionInState,
				TuitionOutOfState: r.TuitionOutOfState,
			}
			dataOut = append(dataOut, school)
		}
		page++
	}
	return dataOut
}
