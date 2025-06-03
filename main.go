package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
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

func handle(conn net.Conn) {
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

		switch input {
		case "1":
			cmd := exec.Command("mdcat", "./pages/about.md")
			output, _ := cmd.Output()
			conn.Write(output)
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
	ln, err := net.Listen("tcp", ":2001")
	if err != nil {
		panic(err)
	}

	log.Println("TCP server running on :2001")

	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}
