package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var htmls = []string{
	"./html/main.html",
	"./html/list.html",
	"./html/edit.html",
	"./html/view.html",
	"./html/session.html",
}

var templates = template.Must(template.ParseFiles(htmls...))

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if len(r.URL.Path) > len("/session/") {
			return
		}
		renderTemplate(w, "session", nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) > len("/list/") {
		return
	}
	p, _ := LoadFiles()
	err := templates.ExecuteTemplate(w, "list.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	if len(r.URL.Path) == len("/view/") {
		http.Redirect(w, r, "/list/", http.StatusFound)
	}

	p, err := LoadFile(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	if len(r.URL.Path) == len("/edit/") {
		return
	}
	p, err := LoadFile(title)
	if err != nil {
		p = &Page{title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	if len(r.URL.Path) == len("/save/") {
		return
	}
	body := r.FormValue("body")
	p := &Page{title: title, body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
