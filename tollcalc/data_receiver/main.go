package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/HsiaoCz/code-beast/tollcalc/types"
	"github.com/gorilla/websocket"
)

func main() {
	recv := NewDataReceiver()
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30001", nil)
}

type DataReceiver struct {
	conn  *websocket.Conn
	msgch chan types.OBUData
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
	}
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	upgrade := websocket.Upgrader{
		HandshakeTimeout: time.Second * 10,
		WriteBufferSize:  1024,
		ReadBufferSize:   1024,
	}
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn

	go dr.wsReceiveLoop(r)
}

func (dr *DataReceiver) wsReceiveLoop(r *http.Request) {
	slog.Info("client connected", "remoteAddr", r.RemoteAddr)
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			slog.Error("data receiver read from websocket error", "err", err)
			continue
		}
		fmt.Printf("received OBU data from [%d] :: <lat %2.f,long %2.f>\n", data.OBUID, data.Lat, data.Long)
		dr.msgch <- data
	}
}
