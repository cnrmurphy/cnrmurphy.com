package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	FontBold       = "\033[1m"
	FontBlue       = "\033[34m"
	FontResetStyle = "\033[0m"
)

func Bold(s string) string {
	return fmt.Sprintf("%s%s%s", FontBold, s, FontResetStyle)
}

func Blue(s string) string {
	return fmt.Sprintf("%s%s%s", FontBlue, s, FontResetStyle)
}

func ConcatResume() []byte {
	cmd := exec.Command("bash", "-c", "(for f in ./pages/contact.md ./pages/experience.md ./pages/projects.md; do cat \"$f\"; echo -e \"\n\n---\n\n\"; done ) | mdcat")
	output, _ := cmd.Output()
	return output
}

func MakeArticlesListMDFile() error {
	entries, err := os.ReadDir("./pages/articles")
	if err != nil {
		return err
	}

	md := "# Articles\n"

	for _, e := range entries {
		log.Info().Msg(e.Name())
		n := strings.Split(e.Name(), ".")
		md += fmt.Sprintf("* %s\n", n[0])
	}

	mdFile, err := os.OpenFile("./pages/articles_list.md", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open articles_list.md for writing")
	}

	mdFile.WriteString(md)

	return nil
}

func HandleArticles(flags []string) []byte {
	if len(flags) <= 0 {
		return []byte("Please pass an article name to retrieve an article. For a list of articles pass the flag -l. For help, pass the flag -h.")
	}

	flag := flags[0]

	if flag[0] == '-' {
		switch flag {
		case "-l":
			cmd := exec.Command("mdcat", "./pages/articles_list.md")
			output, _ := cmd.Output()
			return output
		default:
			return []byte("command not recognized")
		}
	}

	cmd := exec.Command("mdcat", "./pages/articles/"+flag+".md")
	output, _ := cmd.Output()
	return output
}

func handle(conn net.Conn, l zerolog.Logger) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	conn.Write([]byte(Bold("Welcome to Conor Murphy's server!\n\n")))

	for {
		conn.Write([]byte(Bold("→ ")))
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		input := strings.TrimSpace(line)
		input = strings.ToLower(input)
		parts := strings.Split(input, " ")

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		l.Info().Msg(input)

		switch command {
		case "about":
			cmd := exec.Command("mdcat", "./pages/about.md")
			output, _ := cmd.Output()
			conn.Write(output)
		case "articles":
			o := HandleArticles(parts[1:])
			conn.Write(o)
		case "contact":
			cmd := exec.Command("mdcat", "./pages/contact.md")
			output, _ := cmd.Output()
			conn.Write(output)
		case "experience":
			cmd := exec.Command("mdcat", "./pages/experience.md")
			output, _ := cmd.Output()
			conn.Write(output)
		case "projects":
			cmd := exec.Command("mdcat", "./pages/projects.md")
			output, _ := cmd.Output()
			conn.Write(output)
		case "resume":
			output := ConcatResume()
			conn.Write(output)
		case "bye", "quit", "exit":
			conn.Write([]byte(Bold("Thanks for visiting, take care!\n")))
			return
		default:
			conn.Write([]byte("Command not recognized!\n"))
		}
	}
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open log file")
	}

	log.Logger = zerolog.New(logFile).With().Timestamp().Logger()

	MakeArticlesListMDFile()

	log.Info().Msg("starting TCP server on :2001")

	ln, err := net.Listen("tcp", ":2001")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen on :2001")
	}

	log.Info().Msg("TCP server running on :2001")

	for {
		conn, _ := ln.Accept()
		l := log.With().Str("remoteAddr", conn.RemoteAddr().String()).Logger()
		l.Info().Msg("connection accepted")
		go handle(conn, l)
	}
}
