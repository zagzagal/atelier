package main

import (
	ad "bitbucket.org/zagzagal/AtelierComplete/AtelierData"
	"log"
	"net/http"
)

var DATA *ad.AtelierData

func webui(filePath string) {
	if filePath == "" {
		DATA = ad.NewAtelier()
	} else {
		DATA = loadData(filePath)
	}

	router := NewRouter()
	log.Printf("Serving on 8080 from file %s\n", filePath)

	log.Fatal(http.ListenAndServe(":8080", router))
}
