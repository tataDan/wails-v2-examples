package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type OptionsType struct {
	Small     bool
	Medium    bool
	Large     bool
	Thin      bool
	Thick     bool
	Pepperoni bool
	Sausage   bool
	Linguica  bool
	Anchovies bool
	Mushrooms bool
	Olives    bool
	Tomatoes  bool
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

func (a *App) Order(name string, phone string, address string, options OptionsType) string {
	var size_options string
	var crust_options string

	if options.Small {
		size_options = "small"
	} else if options.Medium {
		size_options = "medium"
	} else if options.Large {
		size_options = "large"
	}

	if options.Thin {
		crust_options = "thin"
	} else if options.Thick {
		crust_options = "thick"
	}

	topping_options := ""
	num_toppings := 0
	len_last_topping := 0

	if options.Pepperoni {
		if num_toppings > 0 {
			topping_options = fmt.Sprintf("%s%s", topping_options, ", ")
		}
		topping_to_add := "pepperoni"
		topping_options = fmt.Sprintf("%s %s", topping_options, topping_to_add)
		len_last_topping = len(topping_to_add)
		num_toppings += 1
	}

	if options.Sausage {
		if num_toppings > 0 {
			topping_options = fmt.Sprintf("%s%s", topping_options, ", ")
		}
		topping_to_add := "sausage"
		topping_options = fmt.Sprintf("%s%s", topping_options, topping_to_add)
		len_last_topping = len(topping_to_add)
		num_toppings += 1
	}

	if options.Linguica {
		if num_toppings > 0 {
			topping_options = fmt.Sprintf("%s%s", topping_options, ", ")
		}
		topping_to_add := "linguica"
		topping_options = fmt.Sprintf("%s %s", topping_options, topping_to_add)
		len_last_topping = len(topping_to_add)
		num_toppings += 1
	}

	if options.Anchovies {
		if num_toppings > 0 {
			topping_options = fmt.Sprintf("%s%s", topping_options, ", ")
		}
		topping_to_add := "anchovies"
		topping_options = fmt.Sprintf("%s %s", topping_options, topping_to_add)
		len_last_topping = len(topping_to_add)
		num_toppings += 1
	}

	if options.Olives {
		if num_toppings > 0 {
			topping_options = fmt.Sprintf("%s%s", topping_options, ", ")
		}
		topping_to_add := "olives"
		topping_options = fmt.Sprintf("%s%s", topping_options, topping_to_add)
		len_last_topping = len(topping_to_add)
		num_toppings += 1
	}

	if options.Mushrooms {
		if num_toppings > 0 {
			topping_options = fmt.Sprintf("%s%s", topping_options, ", ")
		}
		topping_to_add := "mushrooms"
		topping_options = fmt.Sprintf("%s%s", topping_options, topping_to_add)
		len_last_topping = len(topping_to_add)
		num_toppings += 1
	}

	if options.Tomatoes {
		if num_toppings > 0 {
			topping_options = fmt.Sprintf("%s%s", topping_options, ", ")
		}
		topping_to_add := "tomatoes"
		topping_options = fmt.Sprintf("%s %s", topping_options, topping_to_add)
		len_last_topping = len(topping_to_add)
		num_toppings += 1
	}

	if num_toppings == 0 {
		topping_options = ""
	} else if num_toppings > 1 {
		topping_options = topping_options[:len(topping_options)-len_last_topping] + "and " + topping_options[len(topping_options)-len_last_topping:]
	}
	if num_toppings == 2 {
		topping_options = strings.ReplaceAll(topping_options, ",", "")
	}
	var with_str string
	if len(topping_options) == 0 {
		with_str = ""
	} else {
		with_str = " with "
	}

	order := fmt.Sprintf("%s\n%s\n%s\n\nYour ordered a %s, %s-crust pizza%s%s",
		name, phone, address, size_options, crust_options, with_str, topping_options)

	_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Title:         "YOUR PIZZA ORDER!",
		Message:       order,
		Buttons:       []string{"OK"},
		DefaultButton: "OK",
	})

	if err != nil {
		log.Println(err)
	}

	return order
}
