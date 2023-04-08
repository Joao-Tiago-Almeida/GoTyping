package terminal

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
)

// Defining colors
var GREEN = color.New(color.FgGreen).SprintFunc()
var RED = color.New(color.FgRed).SprintFunc()
var WHITE = color.New(color.FgWhite).SprintFunc()

func TerminalBox(ch_lessons chan []string) {

	init_termbox()

	for {
		select {
		case message := <-ch_lessons:
			if message[0] == "terminal" {
				if interpret_message(message[1]) {
					ch_lessons <- []string{"lessons", "yes"}
				} else {
					ch_lessons <- []string{"lessons", "no"}
					close_termbox()
					return
				}
			} else {
				ch_lessons <- message // send the message back to the channel
			}

		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

// Might do something with this later, for now it just a new dummy message
func interpret_message(message string) bool {
	done := run_lesson(message)
	return done
}

func init_termbox() {
	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func close_termbox() {
	termbox.Close()
}

func run_lesson(sentence string) bool {

	n_words := count_words(sentence, "_")
	current_key := ""
	n_keys_right := 0
	running := false
	var start time.Time
	var elapsed float64

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
	fmt.Printf("%s", WHITE(sentence))
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
				return false
			} else if event.Ch != 0 {
				// Write the pressed key to the channel
				current_key = string(event.Ch)
			} else if event.Key == termbox.KeySpace {
				// Handle space
				current_key = "_"
			}
			n_keys_right = update_key(current_key, sentence, n_keys_right)

			// Check if it is the last sentence
			if n_keys_right == len(sentence) {
				elapsed = float64(time.Since(start).Milliseconds()) / 1000.0
				running = false
				fmt.Printf(" :: wpm: %.3f\n", 60.0*float64(n_words)/elapsed)
				return true
			}

		default:
			// No events, so do some other work here
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func count_words(s string, del string) int {
	count := 1

	for _, c := range s {
		if string(c) == del {
			count++
		}
	}
	return count
}

func update_key(key string, sentence string, n_keys_right int) int {

	var correct_keys, incorrect_key, remaining_keys string = "", "", ""

	if key == string(sentence[n_keys_right]) {
		n_keys_right++ // increment the number of correct keys
		correct_keys = sentence[:n_keys_right]
		remaining_keys = sentence[n_keys_right:]
	} else {
		correct_keys = sentence[:n_keys_right]
		incorrect_key = string(sentence[n_keys_right])
		remaining_keys = sentence[n_keys_right+1:]
	}

	fmt.Printf("\r%s", GREEN(correct_keys)) // already correct keys
	fmt.Printf("%s", RED(incorrect_key))    // current incorrect  key
	fmt.Printf("%s", WHITE(remaining_keys)) // remaining keys
	return n_keys_right
}
