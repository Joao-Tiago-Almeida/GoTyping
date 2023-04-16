package lessons

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func Wikipedia(title string) []string {

	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&titles=%s&prop=extracts&format=json&exsectionformat=plain", title)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Extract plain text from the response body
	plainText := extractPlainText(body)
	// fmt.Println(plainText)
	return plainText
}

func extractPlainText(body []byte) []string {
	// Parse the response body and extract the plain text
	// You can use a JSON parsing library to extract the plain text from the response body
	// Here, we use a simple string manipulation to extract the plain text
	start := "\"extract\":\""
	end := "\"}}"
	indexStart := len(start)
	indexEnd := len(body) - len(end)

	extractedText := string(body[indexStart:indexEnd])

	// Create a regular expression to match <p> and </p> tags and the text between them
	re := regexp.MustCompile(`<p>(.*?)<\/p>`)

	// Find all the matches of the regular expression in the text
	matches := re.FindAllStringSubmatch(extractedText, -1)

	// Create the corpus from the matches the second element of the matches is the text between the <p> and </p> tags
	corpus := make([]string, len(matches))
	for i, match := range matches {
		corpus[i] = process_text(match[1])
	}

	return corpus
}

func process_text(input_text string) string {

	// Remove all non alphanumeric characters and change all characters to lower case
	re := regexp.MustCompile(`[^a-zA-Z ]+`)
	processed_text := re.ReplaceAllString(input_text, "")
	processed_text = strings.ToLower(processed_text)
	return processed_text
}
