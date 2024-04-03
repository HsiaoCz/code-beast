package main

import (
	"testing"
	"time"
)

func FetchResource() string {
	time.Sleep(time.Second * 2)
	return "some result"
}

func TestFetchResource(t *testing.T) {
	start := time.Now()
	FetchResource()
	FetchResource()
	FetchResource()
	FetchResource()
	FetchResource()
	t.Log("the time use:", time.Since(start))
}
