package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/HsiaoCz/code-beast/tollcalc/types"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/websocket"
)

var kafkaTopic = "obudata"

func main() {

	// p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	// if err != nil {
	// 	panic(err)
	// }

	// defer p.Close()

	// // Delivery report handler for produced messages
	// go func() {
	// 	for e := range p.Events() {
	// 		switch ev := e.(type) {
	// 		case *kafka.Message:
	// 			if ev.TopicPartition.Error != nil {
	// 				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	// 			} else {
	// 				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
	// 			}
	// 		}
	// 	}
	// }()

	// // Produce messages to topic (asynchronously)
	// topic := kafkaTopic
	// p.Produce(&kafka.Message{
	// 	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 	Value:          []byte("hello my man"),
	// }, nil)
	// Wait for message deliveries before shutting down
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30001", nil)
}

type DataReceiver struct {
	conn  *websocket.Conn
	msgch chan types.OBUData
	prod  *kafka.Producer
}

func NewDataReceiver() (*DataReceiver, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"": ""})
	if err != nil {
		return nil, err
	}
	// start another goroutine to check if we have delivered the data
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
		prod:  p,
	}, nil
}

func (dr *DataReceiver) produceData(data types.OBUData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	dr.prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &kafkaTopic,
			Partition: kafka.PartitionAny,
		},
		Value: b,
	}, nil)
	return nil
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
		if err := dr.produceData(data); err != nil {
			slog.Error("kafka produce data error", "err", err)
			return
		}
	}
}
