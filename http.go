package main

import (
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/rs/zerolog"
)

type HTTPServer struct {
	Logger zerolog.Logger
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

func (s *HTTPServer) Start() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/resume", ResumeHandler)
	http.HandleFunc("/articles", ArticlesHandler)
	http.HandleFunc("/articles/", ArticlesHandler)
	http.HandleFunc("/contact", ContactHandler)
	http.ListenAndServe(":8080", nil)
}

type Page struct {
	Title   string
	Content []byte
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	c, _ := os.ReadFile("./dist/about.html")
	t, _ := template.ParseFiles("./wrapper.html")

	p := &Page{
		Title:   "Conor Murphy",
		Content: c,
	}

	t.Execute(w, p)
}

func ContactHandler(w http.ResponseWriter, req *http.Request) {
	c, _ := os.ReadFile("./dist/contact.html")
	t, _ := template.ParseFiles("./wrapper.html")

	p := &Page{
		Title:   "Conor Murphy",
		Content: c,
	}

	t.Execute(w, p)
}

func ResumeHandler(w http.ResponseWriter, req *http.Request) {
	c, _ := os.ReadFile("./dist/resume.html")
	t, _ := template.ParseFiles("./wrapper.html")

	p := &Page{
		Title:   "Resume",
		Content: c,
	}

	t.Execute(w, p)
}

func ArticlesHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	p := &Page{
		Title: "Articles",
	}
	t, _ := template.ParseFiles("./wrapper.html")

	if path == "/articles" {
		c, _ := os.ReadFile("./dist/articles_list.html")
		p.Content = c
		t.Execute(w, p)
		return
	}

	if strings.HasPrefix(path, "/articles/") {
		articleName := strings.TrimPrefix(path, "/articles/")
		c, _ := os.ReadFile("./dist/articles/" + articleName + ".html")

		p.Title = articleName
		p.Content = c
		t.Execute(w, p)
		return
	}
	http.NotFound(w, req)
}
