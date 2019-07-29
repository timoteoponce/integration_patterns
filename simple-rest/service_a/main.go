package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting our SERVICE-A")
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/documents", delegateRequest)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func delegateRequest(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:9100/documents")
	if err != nil {
		log.Printf("Error on connection %v\n", err)
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(resp.Body)
	w.Write(body)
}
