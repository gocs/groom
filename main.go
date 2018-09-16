package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	LoadFiles()
	fmt.Println("Started...")
	http.HandleFunc("/list/", listHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/session/", sessionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
