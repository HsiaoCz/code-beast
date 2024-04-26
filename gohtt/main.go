package main

import (
	"log"
	"net/http"

	"github.com/HsiaoCz/code-beast/gohtt/view"
	"github.com/HsiaoCz/code-beast/gohtt/view/layout"
	"github.com/HsiaoCz/code-beast/gohtt/view/partial"
	"github.com/a-h/templ"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	c := layout.Base(view.Index())
	http.Handle("/", templ.Handler(c))

	http.Handle("/foo", templ.Handler(partial.Foo()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
