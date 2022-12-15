package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

type Response struct {
	NumFound int
	Docs     []Doc
}

type Doc struct {
	Key         string
	Title       string
	Subtitle    string
	Author_name []string
	Publisher   []string
	Person      []string
	Place       []string
	Subject     []string
	Isbn        []string
	Language    []string
}

type QueryParams struct {
	Title     string
	Author    string
	Publisher string
	Person    string
	Place     string
	Subject   string
	Isbn      string
	Language  string
}

const exactSearch = "exactSearch"
const subStringSearch = "subStringSearch"
const normalSearch = "normalSearch"
const exactTitleSearch = "exactTitleSearch"
const subStringTitleSearch = "subStringTitleSearch"
const exactAuthorSearch = "exactAuthorSearch"
const subStringAuthorSearch = "subStringAuthorSearch"
const exactPersonSearch = "exactPersonSearch"
const subStringPersonSearch = "subStringPersonSearch"
const exactPlaceSearch = "exactPlaceSearch"
const subStringPlaceSearch = "subStringPlaceSearch"
const exactSubjectSearch = "exactSubjectSearch"
const subStringSubjectSearch = "subStringSubjectSearch"
const exactPublisherSearch = "exactPublisherSearch"
const subStringPublisherSearch = "subStringPublisherSearch"

func configQuery(query string) string {
	query = strings.TrimSpace(query)
	return strings.ReplaceAll(query, " ", "+")
}

func configSearch(slice []string, param string, searchType string) bool {
	docAdded := false
	for _, val := range slice {
		if docAdded {
			break
		}
		if searchType == exactSearch {
			if strings.EqualFold(strings.ToUpper(val), strings.ToUpper(strings.TrimSpace(param))) {
				docAdded = true
				return true
			}
		} else if searchType == subStringSearch {
			if len(param) > 0 && (strings.Contains(strings.ToUpper(val), strings.ToUpper(strings.TrimSpace(param)))) {
				docAdded = true
				return true
			}
		}
	}
	return false
}

func (a *App) SearchDocuments(queryParams QueryParams, searchType string) []Doc {
	const docsShownLimit = 1000 // Actual limit could be as high as 1100
	const docsLimit = 2000
	numDocsShown := 0
	var docs = []Doc{}
	urlBase := "https://openlibrary.org/search.json?"
	query := ""
	titleNew := configQuery(queryParams.Title)
	if len(titleNew) > 0 {
		query = "title=" + titleNew
	}
	authorNew := configQuery(queryParams.Author)
	if len(authorNew) > 0 {
		query = query + "&author=" + authorNew
	}
	publisherNew := configQuery(queryParams.Publisher)
	if len(publisherNew) > 0 {
		query = query + "&publisher=" + publisherNew
	}
	personNew := configQuery(queryParams.Person)
	if len(personNew) > 0 {
		query = query + "&person=" + personNew
	}
	placeNew := configQuery(queryParams.Place)
	if len(placeNew) > 0 {
		query = query + "&place=" + placeNew
	}
	subjectNew := configQuery(queryParams.Subject)
	if len(subjectNew) > 0 {
		query = query + "&subject=" + subjectNew
	}
	isbnNew := configQuery(queryParams.Isbn)
	if len(isbnNew) > 0 {
		query = query + "&isbn=" + isbnNew
	}
	languageNew := configQuery(queryParams.Language)
	if len(languageNew) > 0 {
		query = query + "&q=language:" + languageNew
	}
	if len(query) > 0 && (query[0] == 38) { // 38 = &
		query = query[1:]
	}
	warningShown := false
	page := 1
	for {
		if numDocsShown > docsShownLimit {
			break
		}
		url := fmt.Sprintf("%s%s&page=%d", urlBase, query, page)
		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err.Error())
		}
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		var responseObject Response
		json.Unmarshal(responseData, &responseObject)
		if len(responseObject.Docs) == 0 {
			break
		}
		if responseObject.NumFound > docsLimit && !warningShown {
			warningShown = true
			errorMsg := fmt.Sprintf("The number of documents found by the openlibrary web site was %d, which exceeds the limit for this application. A maximum of %d documents will be shown.",
				responseObject.NumFound, (docsShownLimit + 100))
			_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
				Title:         "WARNING",
				Message:       errorMsg,
				Buttons:       []string{"OK"},
				DefaultButton: "OK",
			})
			if err != nil {
				log.Println(err)
				return nil
			}
		}
		page++
		for _, doc := range responseObject.Docs {
			switch searchType {
			case normalSearch:
				docs = append(docs, doc)
			case exactTitleSearch:
				if strings.EqualFold(strings.ToUpper(doc.Title), strings.ToUpper(strings.TrimSpace(queryParams.Title))) {
					docs = append(docs, doc)
				}
			case subStringTitleSearch:
				if strings.Contains(strings.ToUpper(doc.Title), strings.ToUpper(strings.TrimSpace(queryParams.Title))) {
					docs = append(docs, doc)
				}
			case exactAuthorSearch:
				if configSearch(doc.Author_name, queryParams.Author, exactSearch) {
					docs = append(docs, doc)
				}
			case subStringAuthorSearch:
				if configSearch(doc.Author_name, queryParams.Author, subStringSearch) {
					docs = append(docs, doc)
				}
			case exactPersonSearch:
				if configSearch(doc.Person, queryParams.Person, exactSearch) {
					docs = append(docs, doc)
				}
			case subStringPersonSearch:
				if configSearch(doc.Person, queryParams.Person, subStringSearch) {
					docs = append(docs, doc)
				}
			case exactPlaceSearch:
				if configSearch(doc.Place, queryParams.Place, exactSearch) {
					docs = append(docs, doc)
				}
			case subStringPlaceSearch:
				if configSearch(doc.Place, queryParams.Place, subStringSearch) {
					docs = append(docs, doc)
				}
			case exactSubjectSearch:
				if configSearch(doc.Subject, queryParams.Subject, exactSearch) {
					docs = append(docs, doc)
				}
			case subStringSubjectSearch:
				if configSearch(doc.Subject, queryParams.Subject, subStringSearch) {
					docs = append(docs, doc)
				}
			case exactPublisherSearch:
				if configSearch(doc.Publisher, queryParams.Publisher, exactSearch) {
					docs = append(docs, doc)
				}
			case subStringPublisherSearch:
				if configSearch(doc.Publisher, queryParams.Publisher, subStringSearch) {
					docs = append(docs, doc)
				}
			}
			numDocsShown++
		}
	}
	return docs
}
