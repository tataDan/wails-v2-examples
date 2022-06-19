package main

import (
	"context"
	"fmt"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Uppercase returns a uppercase equivalent for the given value
func (a *App) Uppercase(value string) string {
	return fmt.Sprintf("The new value is %s", strings.ToUpper(value))
}
