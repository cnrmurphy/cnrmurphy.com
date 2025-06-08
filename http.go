package main

import (
	"net/http"
	"os/exec"
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
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/resume", ResumeHandler)
	http.HandleFunc("/articles", ArticlesHandler)
	http.ListenAndServe(":8080", nil)
}

type Page struct {
	Title   string
	Content []byte
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("pandoc", "./pages/contact.md", "-f", "markdown", "-t", "html")
	output, _ := cmd.Output()
	t, _ := template.ParseFiles("./wrapper.html")

	p := &Page{
		Title:   "Home",
		Content: output,
	}

	t.Execute(w, p)
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("pandoc", "./pages/about.md", "-f", "markdown", "-t", "html")
	output, _ := cmd.Output()
	t, _ := template.ParseFiles("./wrapper.html")

	p := &Page{
		Title:   "About",
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
	cmd := exec.Command("pandoc", "./pages/articles_list.md", "-f", "markdown", "-t", "html")
	output, _ := cmd.Output()
	t, _ := template.ParseFiles("./wrapper.html")

	p := &Page{
		Title:   "Articles",
		Content: output,
	}

	t.Execute(w, p)
}
