package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://api-gw.onebound.cn/douyin/user_info/?key=t5971078584&secret=20240506&sec_uid=MS4wLjABAAAAKLbzdLrJxBLWIhuaDRJQYGV0sa7xmKvOovj_N0mRPhA"
	apiKey := "t5971078584"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
