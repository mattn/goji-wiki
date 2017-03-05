package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/flosch/pongo2"
	_ "github.com/flosch/pongo2-addons"

	_ "github.com/mattn/go-sqlite3"

	"github.com/naoina/genmai"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

type Wiki struct {
	URL string
	DB  *genmai.DB
}

// Page store content for wiki page
type Page struct {
	Id        int64     `db:"pk"`
	Title     string    `db:"unique" json:"title"`
	Body      string    `json:"body"`
	URL       string    `db:"-"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (page *Page) BeforeInsert() error {
	n := time.Now()
	page.CreatedAt = n
	page.UpdatedAt = n
	return nil
}

func (page *Page) BeforeUpdate() error {
	n := time.Now()
	page.UpdatedAt = n
	return nil
}

func (wiki *Wiki) PageURL(page *Page) string {
	return (&url.URL{Path: path.Join(wiki.URL, "/wiki", page.Title)}).Path
}

func getWiki(c web.C) *Wiki {
	return c.Env["Wiki"].(*Wiki)
}

func main() {
	// setup tables
	db, err := genmai.New(&genmai.SQLite3Dialect{}, "./wiki.db")
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.CreateTableIfNotExists(&Page{}); err != nil {
		log.Fatalln(err)
	}

	// setup pongo
	pongo2.DefaultLoader.SetBaseDir("view")

	wiki := &Wiki{
		URL: "/",
		DB:  db,
	}
	pongo2.Globals["wiki"] = wiki
	pongo2.RegisterFilter("to_localdate", func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
		date, ok := in.Interface().(time.Time)
		if !ok {
			return nil, &pongo2.Error{
				Sender:    "to_localdate",
				OrigError: fmt.Errorf("Date must be of type time.Time not %T ('%v')", in, in),
			}
		}
		return pongo2.AsValue(date.Local()), nil
	})

	goji.Use(middleware.Recoverer)
	goji.Use(middleware.NoCache)
	goji.Use(func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			c.Env["Wiki"] = wiki
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)

	})

	goji.Get("/assets/*", http.FileServer(http.Dir(".")))
	goji.Get("/", showPages)
	goji.Get("/wiki/:title", showPage)
	goji.Get("/wiki/:title/edit", editPage)
	goji.Post("/wiki/:title", postPage)

	goji.Serve()
}
