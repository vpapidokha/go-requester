package requester

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func UrlRequester(url string) {
	defer log.Println("Request sent...")
	response, error := http.Get(url)
	if error != nil {
		log.Println(error)
	}
	log.Println(response)

	n := rand.Intn(10)
	log.Printf("Sleeping %d seconds...\n", n)
	time.Sleep(time.Duration(n) * time.Second)
}
