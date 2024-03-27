package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       string
	Comments []string
	Likes    int
	Friends  []int
}

type Response struct {
	data any
	err  error
}

func handleGetUserProfile(id string) (*UserProfile, error) {
	respch := make(chan Response, 3)
	wg := &sync.WaitGroup{}
	go getComments(id, respch, wg)
	go getLikes(id, respch, wg)
	go getFriends(id, respch, wg)

	wg.Add(3)
	wg.Wait()
	close(respch)
	userProfile := &UserProfile{}

	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}
	}

	return userProfile, nil
}

func getComments(id string, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{
		"Hey,that was great",
		"Yeah buddy",
		"Ow,I didnt know that",
	}
	respch <- Response{
		data: comments,
		err:  nil,
	}
	wg.Done()
}

func getLikes(id string, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	respch <- Response{
		data: 11,
		err:  nil,
	}
	wg.Done()
}

func getFriends(id string, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	respch <- Response{
		data: []int{11, 22, 33, 44, 55, 66},
		err:  nil,
	}
	wg.Done()
}

func main() {
	userProfile, err := handleGetUserProfile("100")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the user profile %+v\n", userProfile)
}
