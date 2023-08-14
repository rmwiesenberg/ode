package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var baseURL = "https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=%s"

func Verber(phrase string) string {
	splitMessage := strings.Split(phrase, " ")
	words := []string{splitMessage[0] + "es", splitMessage[0] + "s"}
	log.Printf("Verbing %s and %s", words[0], words[1])
	for _, word := range words {
		res, err := http.Get(fmt.Sprintf(baseURL, word, os.Getenv("DICTIONARY_API_KEY")))
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			if strings.Contains(bodyString, "\"fl\":\"verb\"") {
				return fmt.Sprintf("%s %s", word, strings.Join(splitMessage[1:], " "))
			}
		}
	}
	return ""
}
