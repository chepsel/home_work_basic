package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	req, err := http.Get("http://localhost:10001/502")
	if err != nil {
		log.Println(err)
	}
	defer req.Body.Close()

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println("localhost:10001 has", string(bodyBytes))

}
