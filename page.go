package main

import (
	"io/ioutil"
	"log"
)

// Page page something
type Page struct {
	title string // title of the page
	body  []byte // body of the page
}

// Title p.title getter
func (p Page) Title() string {
	return p.title
}

// Body p.body getter
func (p Page) Body() []byte {
	return p.body
}

// BodyStr p.body getter string
func (p Page) BodyStr() string {
	return string(p.body)
}

// SetTitle sets title
func (p *Page) SetTitle(title string) {
	p.title = title
}

// SetBody sets body
func (p *Page) SetBody(body []byte) {
	p.body = body
}

// Save saves file
func (p *Page) Save() error {
	filename := "./text/" + p.title + ".txt"
	return ioutil.WriteFile(filename, p.body, 0600)
}

// LoadFile loads file
func LoadFile(title string) (*Page, error) {
	filename := "./text/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title: title, body: body}, nil
}

// LoadFiles loads files
func LoadFiles() ([]Page, error) {
	files, err := ioutil.ReadDir("./text/")
	if err != nil {
		log.Fatal(err)
	}

	var pages []Page
	for _, f := range files {
		filename := "./text/" + f.Name()
		body, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		pages = append(pages, Page{title: f.Name()[:len(f.Name())-4], body: body})
	}
	return pages, nil
}
