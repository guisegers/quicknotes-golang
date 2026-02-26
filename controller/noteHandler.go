package controller

import (
	"net/http"
	"text/template"
)

// NoteHandler implements http.Handler for the note list (root path).
type NoteHandler struct{}

func NoteList(w http.ResponseWriter, r *http.Request) {
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

func NoteView(w http.ResponseWriter, r *http.Request) {
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

func NoteCreate(w http.ResponseWriter, r *http.Request) {
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
