package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
)

func countWords(s string) int {
	count := 1

	for _, c := range s {
		if c == ' ' {
			count++
		}
	}
	return count
}

func main() {

	// Define the sentence to be guessed
	sentence := "This is a sentence to be guessed."
	n_words := countWords(sentence)
	current_key := ""
	n_keys_right := 0
	running := false
	var start time.Time
	var elapsed float64

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	// defer termbox.Close()

	// Create a channel to receive keyboard events
	keyEvents := make(chan termbox.Event)

	// Start a goroutine to listen for keyboard events and write them to the channel
	go func() {
		for {
			event := termbox.PollEvent()
			if event.Type == termbox.EventKey {
				keyEvents <- event
			}
		}
	}()

	// Main loop
	fmt.Println(sentence)
	for {
		select {
		case event := <-keyEvents:
			// if it is the first key pressed, start the timer
			if !running {
				start = time.Now()
				running = true
			}

			// Handle the key event
			if event.Key == termbox.KeyEsc {
				// Exit the program if the user presses the Esc key
				return
			} else if event.Ch != 0 {
				// Write the pressed key to the channel
				current_key = string(event.Ch)
			} else if event.Key == termbox.KeySpace {
				// Handle space
				current_key = " "
			}
			if current_key == string(sentence[n_keys_right]) {
				n_keys_right++
				// Check if it is the last sentence
				if n_keys_right == len(sentence) {
					elapsed = float64(time.Since(start).Milliseconds()) / 1000.0
					running = false
					termbox.Close()
					fmt.Printf("wpm: %.3f\n", 60.0*float64(n_words)/elapsed)
					return
				}
				fmt.Printf("\r%s", green(sentence[:n_keys_right]))
			} else {
				fmt.Printf("\r%s", red(sentence[:n_keys_right]+current_key))
			}

		default:
			// No events, so do some other work here
			time.Sleep(time.Millisecond * 1)
		}
	}
}
