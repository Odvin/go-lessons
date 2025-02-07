package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/template"
)

type Todo struct {
	ID   int
	Task string
	Done bool
}

var (
	todos     []Todo
	idCounter int
	mutex     sync.Mutex

	indexTemplate *template.Template
)

type IndexData struct {
	Todos []TodoView
}

type TodoView struct {
	ID           int
	Task         string
	Done         bool
	CheckedClass string
	CheckIcon    string
	TextStyle    string
}

func main() {
	var err error
	indexTemplate, err = template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Println("Listening on port 3000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	reader := bufio.NewReader(conn)
	request, _ := reader.ReadString('\n')
	if len(request) == 0 {
		return
	}

	parts := strings.Fields(request)
	if len(parts) < 2 {
		return
	}

	method, path := parts[0], parts[1]

	switch method {
	case "GET":
		handleGet(conn, path)
	case "POST":
		handlePost(conn, path, reader)
	default:
		sendResponse(conn, 400, "Bad Request")
	}
}

func handlePost(conn net.Conn, path string, reader *bufio.Reader) {
	var contentLength int

	for {
		line, _ := reader.ReadString('\n')
		if line == "\r\n" {
			break
		}

		if strings.HasPrefix(line, "Content-Length") {
			contentLength, _ = strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1]))
		}
	}

	body := make([]byte, contentLength)
	_, err := reader.Read(body)
	if err != nil {
		sendResponse(conn, 500, "Internal Server Error")
		return
	}

	switch path {
	case "/add":
		handleAdd(conn, body)
	case "/delete":
		handleDelete(conn, body)
	case "/toggle":
		handleToggle(conn, body)
	default:
		sendResponse(conn, 404, "Not Found")
	}
}

func handleToggle(conn net.Conn, body []byte) {

	params := parseForm(string(body))
	id, _ := strconv.Atoi(params["id"])

	mutex.Lock()
	for i, t := range todos {
		if t.ID == id {
			todos[i].Done = !todos[i].Done
			break
		}
	}

	mutex.Unlock()
	sendRedirect(conn, "/")

}

func handleDelete(conn net.Conn, body []byte) {

	params := parseForm(string(body))
	id, _ := strconv.Atoi(params["id"])

	mutex.Lock()
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

	mutex.Unlock()
	sendRedirect(conn, "/")
}

func handleAdd(conn net.Conn, body []byte) {
	params := parseForm(string(body))
	task := params["task"]

	if task != "" {
		mutex.Lock()
		idCounter++
		todos = append(todos, Todo{
			ID:   idCounter,
			Task: task,
		})
		mutex.Unlock()
	}

	sendRedirect(conn, "/")
}

func parseForm(data string) map[string]string {
	params := make(map[string]string)
	pairs := strings.Split(data, "&")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			params[kv[0]] = strings.ReplaceAll(kv[1], "+", " ")
		}
	}
	return params
}

func handleGet(conn net.Conn, path string) {
	if path != "/" {
		sendResponse(conn, 404, "Not Found")
	}

	mutex.Lock()
	defer mutex.Unlock()

	views := make([]TodoView, 0, len(todos))

	for _, t := range todos {
		view := TodoView{
			ID:   t.ID,
			Task: t.Task,
			Done: t.Done,
		}

		if t.Done {
			view.CheckedClass = "bg-blue-500 border-blue-500"
			view.CheckIcon = "✓"
			view.TextStyle = "line-through text-gray-400"
		} else {
			view.CheckedClass = "border-gray-300"
			view.TextStyle = "text-gray-700"
		}

		views = append(views, view)
	}

	data := IndexData{
		Todos: views,
	}

	var pageBuf bytes.Buffer
	err := indexTemplate.Execute(&pageBuf, data)
	if err != nil {
		sendResponse(conn, 500, "Internal Server Error")
		return
	}

	sendResponse(conn, 200, pageBuf.String())
}

func sendResponse(conn net.Conn, statusCode int, body string) {
	response := fmt.Sprintf("HTTP/1.1 %d OK\r\n", statusCode)
	response += "Content-Type: text/html\r\n"
	response += fmt.Sprintf("Content-Length: %d\r\n", len(body))
	response += "\r\n"
	response += body

	conn.Write([]byte(response))
	conn.Close()
}

func sendRedirect(conn net.Conn, path string) {
	response := "HTTP/1.1 302 Found\r\n"
	response += fmt.Sprintf("Location: %s\r\n", path)
	response += "\r\n"
	conn.Write([]byte(response))
	conn.Close()
}
