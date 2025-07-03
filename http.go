package main

import (
	"net/http"
	"os"
	"strings"

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
	c, _ := os.ReadFile("./public/about.html")
	w.Header().Set("Content-Type", "text/html")
	w.Write(c)
}

func ContactHandler(w http.ResponseWriter, req *http.Request) {
	c, _ := os.ReadFile("./public/contact.html")
	w.Header().Set("Content-Type", "text/html")
	w.Write(c)
}

func ResumeHandler(w http.ResponseWriter, req *http.Request) {
	c, _ := os.ReadFile("./public/resume.html")
	w.Header().Set("Content-Type", "text/html")
	w.Write(c)
}

func ArticlesHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if path == "/articles" {
		c, _ := os.ReadFile("./public/articles_list.html")
		w.Header().Set("Content-Type", "text/html")
		w.Write(c)
		return
	}

	if strings.HasPrefix(path, "/articles/") {
		articleName := strings.TrimPrefix(path, "/articles/")
		c, _ := os.ReadFile("./public/articles/" + articleName + ".html")
		w.Header().Set("Content-Type", "text/html")
		w.Write(c)
		return
	}
	http.NotFound(w, req)
}
