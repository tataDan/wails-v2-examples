package main

import (
	"context"
	"math/rand"
	"time"
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

type Point struct {
	X int
	Y int
}

const numPointsMax = 80

func (a *App) SingleSlitExperiment() []Point {
	var points = []Point{}
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for i := 0; i < numPointsMax; i++ {
		point := Point{450 + r.Intn(37), 10 + rand.Intn(277)}
		points = append(points, point)
	}
	return points
}

func (a *App) DoubleSlitExperiment() []Point {
	var points = []Point{}
	var pointsFound = 0

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for pointsFound < numPointsMax {
		point := Point{450 + r.Intn(37), 15 + rand.Intn(277)}
		if point.Y >= 15 && point.Y <= 45 {
			points = append(points, point)
			pointsFound++
		} else if point.Y >= 95 && point.Y <= 125 {
			points = append(points, point)
			pointsFound++
		} else if point.Y >= 175 && point.Y <= 205 {
			points = append(points, point)
			pointsFound++
		} else if point.Y >= 255 && point.Y <= 285 {
			points = append(points, point)
			pointsFound++
		}
	}
	return points
}
