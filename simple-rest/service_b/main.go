package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Document struct {
	ID   string
	Name string
	Size int
}

func main() {
	fmt.Println("Starting our SERVICE-B")
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/documents", getDocuments)
	log.Fatal(http.ListenAndServe(":9100", nil))
}

func getDocuments(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	var docs []Document
	docs = append(docs, Document{ID: "Doc-1", Name: "Report-quarter-1.pdf", Size: 1889})
	docs = append(docs, Document{ID: "Doc-2", Name: "pride-and_prejudice.pdf", Size: 20000})
	docs = append(docs, Document{ID: "Doc-3", Name: "macondo.pdf", Size: 500})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}
