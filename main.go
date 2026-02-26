package main

import (
	"fmt"
	"net/http"

	"quicknotes/controller"
)

func main() {
	fmt.Println("Servidor rodando na porta 8080")
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.NoteList)
	mux.HandleFunc("/note/view", controller.NoteView)
	mux.HandleFunc("/note/create", controller.NoteCreate)

	http.ListenAndServe(":8080", mux)
}
