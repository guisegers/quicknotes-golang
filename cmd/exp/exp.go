package main

import (
	"log"
	"os"
	"text/template"
)

type TemplateData struct {
	Name string
	Age  int
}

func main() {
	t, err := template.ParseFiles("layout1.html", "home.html", "header.html", "footer.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
