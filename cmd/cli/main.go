package main

import (
	"io"
	"log"
	"os"

	"github.com/vpapidokha/go-requester/internal/requester"
)

func main() {
	targetUrl := os.Args[1]

	LOG_FILE := "/var/log/go-requester.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for {
		requester.UrlRequester(targetUrl)
	}
}
