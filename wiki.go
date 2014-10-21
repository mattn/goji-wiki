package main

import (
	"net/http"
	"sync"

	"github.com/flosch/pongo2"
	"github.com/naoina/genmai"
	"github.com/zenazn/goji/web"
)

// editPage display edit box
func editPage(c web.C, w http.ResponseWriter, r *http.Request) {
	wiki := getWiki(c)
	db := wiki.DB

	var pages []Page
	err := db.Select(&pages, db.From(&Page{}), db.Where("title", "=", c.URLParams["title"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var page Page
	if len(pages) > 0 {
		page = pages[0]
	}
	if page.Title == "" {
		page.Title = c.URLParams["title"]
	}
	page.URL = wiki.PageURL(&page)
	tpl, err := pongo2.DefaultSet.FromFile("edit.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteWriter(pongo2.Context{"page": page}, w)
}

// postPage send body content
func postPage(c web.C, w http.ResponseWriter, r *http.Request) {
	wiki := getWiki(c)
	db := wiki.DB

	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	var pages []Page
	err := db.Select(&pages, db.From(&Page{}), db.Where("title", "=", c.URLParams["title"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.FormValue("body") == "" {
		// if body is empty, then delete the page
		if len(pages) > 0 {
			page := pages[0]
			page.Deleted = true
			_, err := db.Update(&page)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// if page exists, update the page
		if len(pages) > 0 {
			page := pages[0]
			page.Body = r.FormValue("body")
			_, err := db.Update(&page)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			page := Page{
				Title: c.URLParams["title"],
				Body:  r.FormValue("body"),
			}
			_, err = db.Insert(&page)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, r.URL.RequestURI(), http.StatusFound)
	}
}

// showPage display the page
func showPage(c web.C, w http.ResponseWriter, r *http.Request) {
	wiki := getWiki(c)
	db := wiki.DB

	var pages []Page
	err := db.Select(&pages, db.From(&Page{}), db.Where("title", "=", c.URLParams["title"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var page Page
	if len(pages) > 0 {
		page = pages[0]
	}
	if page.Title == "" {
		page.Title = c.URLParams["title"]
	}
	page.URL = wiki.PageURL(&page)
	tpl, err := pongo2.DefaultSet.FromFile("page.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteWriter(pongo2.Context{"page": page}, w)
}

// showPages displayes title index
func showPages(c web.C, w http.ResponseWriter, r *http.Request) {
	wiki := getWiki(c)
	db := wiki.DB

	var pages []Page
	err := db.Select(&pages, db.OrderBy("title", genmai.ASC))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl, err := pongo2.DefaultSet.FromFile("index.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i := range pages {
		pages[i].URL = wiki.PageURL(&pages[i])
	}
	tpl.ExecuteWriter(pongo2.Context{"pages": pages}, w)
}
