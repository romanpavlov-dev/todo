package main

import (
	"fmt"
	"todo/http"
	"todo/logics"
)

func main() {
	todolist := logics.NewList()

	HTTPhandlers := http.NewHTTPHandlers(todolist)

	httpServer := http.NewHTTPServer(HTTPhandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Failed to start a server:", err)
	}
}
