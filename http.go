package main

import (
	"net/http"
	"os/exec"
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
	cmd := exec.Command("pandoc", "./pages/about.md", "-f", "markdown", "-t", "html")
	output, _ := cmd.Output()
	t, _ := template.ParseFiles("./wrapper.html")

	p := &Page{
		Title:   "Conor Murphy",
		Content: output,
	}

	t.Execute(w, p)
}

func ContactHandler(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("pandoc", "./pages/contact.md", "-f", "markdown", "-t", "html")
	output, _ := cmd.Output()
	t, _ := template.ParseFiles("./wrapper.html")

	p := &Page{
		Title:   "Conor Murphy",
		Content: output,
	}

	t.Execute(w, p)
}

func ResumeHandler(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("pandoc", "./pages/resume.md", "-f", "markdown", "-t", "html")
	output, _ := cmd.Output()
	t, _ := template.ParseFiles("./wrapper.html")

	p := &Page{
		Title:   "Resume",
		Content: output,
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
		cmd := exec.Command("pandoc", "./pages/articles_list.md", "-f", "markdown", "-t", "html")
		output, _ := cmd.Output()
		p.Content = output
		t.Execute(w, p)
		return
	}

	if strings.HasPrefix(path, "/articles/") {
		articleName := strings.TrimPrefix(path, "/articles/")

		cmd := exec.Command("pandoc", "./pages/articles/"+articleName+".md", "-f", "markdown", "-t", "html")
		output, _ := cmd.Output()
		p.Title = articleName
		p.Content = output
		t.Execute(w, p)
		return
	}
	http.NotFound(w, req)
}
