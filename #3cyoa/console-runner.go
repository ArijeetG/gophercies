package main

import (
	"fmt"
	"log"
	"os"
)

type ConsoleRunner struct {
}

func (cr ConsoleRunner) Start(provider *StoryArcProvider) {
	cr.displayArcText(*provider, "intro")
}

func (cr ConsoleRunner) displayArcText(provider StoryArcProvider, arcName string) {

	arc, err := provider.WriteTemplateText(os.Stdout, arcName)
	if err != nil {
		log.Println(err)
	}

	if len(arc.Options) == 0 {
		return
	}

	fmt.Println("Your Options: ")
	var optionNumber int
	fmt.Scan(&optionNumber)
	for _, option := range arc.Options {
		if option.Number == optionNumber {
			cr.displayArcText(provider, option.Arc)
		}
	}
}
