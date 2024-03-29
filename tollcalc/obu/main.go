package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/HsiaoCz/code-beast/tollcalc/types"
	"github.com/gorilla/websocket"
)

// const sendInterval = time.Second

const wsEndpoint = "ws://127.0.0.1:30001/ws"

// func sendOBUData(data OBUData)error

func genLatLong() (float64, float64) {
	return genCoord(), genCoord()
}

func genCoord() float64 {
	n := float64(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100) + 1)
	f := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
	return n + f
}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(math.MaxInt)
	}
	return ids
}

func main() {
	obuIDS := generateOBUIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(obuIDS); i++ {
			lat, long := genLatLong()
			data := types.OBUData{
				OBUID: obuIDS[i],
				Lat:   lat,
				Long:  long,
			}
			if err := conn.WriteJSON(&data); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println(genCoord())
		// time.Sleep(sendInterval)
	}
}
