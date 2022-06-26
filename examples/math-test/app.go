package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type TestItems struct {
	TestItems []TestItem
}

type TestItem struct {
	Id       int
	Question string
	Answer   string
	Response string
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

func (a *App) TestOne() []TestItem {
	return readJsonFile("test_1.json")
}

func (a *App) TestTwo() []TestItem {
	return readJsonFile("test_2.json")
}

func (a *App) SubmitAnswer(test []TestItem) []bool {
	var num_correct_answers int
	var num_questions int
	var feedback = []bool{}

	for i := 0; i < len(test); i++ {
		num_questions++
		feedback = append(feedback, false)
	}
	for i, t := range test {
		if t.Response == t.Answer {
			num_correct_answers++
			if num_correct_answers == num_questions {
				runtime.EventsEmit(a.ctx, "perfect-score")
			}
			feedback[i] = true
		} else {
			feedback[i] = false
		}
	}

	return feedback
}
