package main

import (
	"github.com/go-typing/src/lessons"
	"github.com/go-typing/src/terminal"
)

func main() {

	// Create a channel to send lessons to the terminal
	ch_lessons := make(chan []string)

	go lessons.GenerateLessons(ch_lessons)
	terminal.TerminalBox(ch_lessons)

}
