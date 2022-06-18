package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Runner interface {
	Start(provider *StoryArcProvider)
}

func main() {
	fmt.Println("Press C to start the console and ENTER otherwise webserver will be started in 5 seconds...")
	keyEntered := make(chan string)
	timeout := time.NewTimer(5 * time.Second)
	go readInputFromUser(keyEntered)
	select {
	case input := <-keyEntered:
		if strings.HasPrefix(strings.ToLower(input), "c") {
			fmt.Println("Console mode selected!!")
			initialiseAndStart(ConsoleRunner{})
		} else {
			fmt.Println("Starting WebRunner")
			initialiseAndStart(WebRunner{})
		}

	case <-timeout.C:
		fmt.Println("Starting WebRunner")
		initialiseAndStart(WebRunner{})

	}
}

func readInputFromUser(keyEntered chan string) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	keyEntered <- input
}

func initialiseAndStart(runner Runner) {
	story := new(Story)
	err := story.Load("gopher.json")
	if err != nil {
		fmt.Println("Stopping ...")
		return
	}

	tt := ConsoleTemplate
	_, ok := runner.(WebRunner)
	if ok {
		tt = WebTemplate
	}

	provider := &StoryArcProvider{
		Story:        story,
		TemplateType: tt,
	}

	err = provider.Initialise()
	if err != nil {
		log.Panicln(err)
		return
	}

	runner.Start(provider)
}
