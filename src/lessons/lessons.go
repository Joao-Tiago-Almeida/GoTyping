package lessons

import (
	"fmt"
	"time"
)

func GenerateLessons(ch_lessons chan []string) {
	// Generate lessons and send them to the channel
	sentence := "This_is_a_sentence_to_be_guessed."
	broadcast_message := []string{"terminal", sentence}

	ch_lessons <- broadcast_message

	for {
		select {
		case message := <-ch_lessons:
			if message[0] == "lessons" {
				if message[1] == "yes" {
					ch_lessons <- broadcast_message
				} else {
					fmt.Println("Nice pratice!")
					close(ch_lessons)
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
