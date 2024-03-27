package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "username", "anthdm")
	result, err := fetchUserID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response took %v --> %+v\n", time.Since(start), result)
}

func ThirdPartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 90)
	return "some response", nil
}

func fetchUserID(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1500)
	defer cancel()

	type result struct {
		userID string
		err    error
	}

	val := ctx.Value("username")
	fmt.Printf("the value is %v", val)

	resultch := make(chan result, 1)

	go func() {
		res, err := ThirdPartyHTTPCall()
		resultch <- result{
			userID: res,
			err:    err,
		}
	}()

	select {
	// Done()
	// 1 --> the context timeout is exceeded
	// 2 --> the context has been manually canceled --> Cancel
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.userID, res.err
	}
}
