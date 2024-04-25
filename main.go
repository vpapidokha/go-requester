package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func requester(url string) {
	defer fmt.Println("Request sent...")
	response, error := http.Get(url)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(response)

	n := rand.Intn(10)
	fmt.Printf("Sleeping %d seconds...\n", n)
	time.Sleep(time.Duration(n) * time.Second)
}

func main() {
	targetUrl := os.Args[1]

	for {
		requester(targetUrl)
	}
}
