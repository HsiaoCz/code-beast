package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/code-beast/ggowg/hollymicro/types"
	"github.com/anthdm/hollywood/actor"
)

const (
	scrapeInterval = time.Second
)

type scraper struct {
	url      string
	storePID *actor.PID
	engine   *actor.Engine
}

func newScraper(url string, storePID *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &scraper{
			url:      url,
			storePID: storePID,
		}
	}
}

func (s *scraper) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		s.engine = c.Engine()
		go s.scrapeLoop()
	case actor.Stopped:
		fmt.Println(msg)
	}
}

func (s *scraper) scrapeLoop() {
	for {
		resp, err := http.Get(s.url)
		if err != nil {
			panic(err)
		}
		var res CatFact
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			log.Println("failed to decode the response body", err)
			continue
		}
		s.engine.Send(s.storePID, &types.CatFact{
			Fact: res.Fact,
		})
		time.Sleep(scrapeInterval)
	}
}

type CatFact struct {
	Fact string `josn:"fact"`
}

func main() {
	listenAddr := flag.String("listenAddr", ":3001", "todo")
	flag.Parse()

	// e, err := actor.NewEngine(actor.EngineConfig{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// r := remote.New(*listenAddr, remote.Config{})

	fmt.Println(listenAddr)

}
