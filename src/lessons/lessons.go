package lessons

import (
	"fmt"
	"math/rand"
	"time"
)

func random_sentence(corpus []string) string {
	return corpus[rand.Int()%len(corpus)]
}

func broadcast_message(message string) []string {
	return []string{"terminal", message}
}

func GenerateLessons(ch_lessons chan []string) {
	// Extract the plain text from the Wikipedia page
	corpus := Wikipedia("Go_(programming_language)")

	// Generate lessons and send them to the channel
	ch_lessons <- broadcast_message(random_sentence(corpus))

	for {
		select {
		case message := <-ch_lessons:
			if message[0] == "lessons" {
				if message[1] == "yes" {
					ch_lessons <- broadcast_message(random_sentence(corpus))
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
