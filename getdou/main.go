package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.douyin.com/user/MS4wLjABAAAAKLbzdLrJxBLWIhuaDRJQYGV0sa7xmKvOovj_N0mRPhA")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
