package main

import (
	"errors"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view|session)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil
}
