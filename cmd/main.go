package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

// NoteHandler implements http.Handler for the note list (root path).
type NoteHandler struct{}

func noteList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
}

func noteView(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
		"views/templates/pages/note-view.html",
	}
	id := r.URL.Query().Get("id")
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", id)
}

func noteNew(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
		"views/templates/pages/note-new.html",
	}
	// if r.Method != http.MethodPost {
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 	return
	// }
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
}

func main() {
	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("views/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/new", noteNew)

	if err := godotenv.Load(); err != nil {
		fmt.Println("Nenhum arquivo .env encontrado, usando variáveis de ambiente do sistema")
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	fmt.Printf("Servidor rodando na porta %s\n", port)

	http.ListenAndServe(":"+port, mux)
}
