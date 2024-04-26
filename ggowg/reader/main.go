package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	rawMsg := []byte("rawbytes")
	handleRawBytes(bytes.NewReader(rawMsg))
	http.HandleFunc("GET /user", handleFoo)
}

// why do we need to use an io.Reader instead of []byte
func handleRawBytes(r io.Reader) {
	http.Post("http://foo.com", "foo", r)
	http.NewRequest("GET", "http://foo.com", r)
}

func handleFoo(w http.ResponseWriter, r *http.Request) {
	herader := w.Header()
	fmt.Println(herader)
	handleRawBytes(r.Body)
}
