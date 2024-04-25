package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func requester(url string) {
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

func main() {
	targetUrl := os.Args[1]

	LOG_FILE := "/var/log/go-requester.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for {
		requester(targetUrl)
	}
}
