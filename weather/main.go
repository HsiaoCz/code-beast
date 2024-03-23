package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	pollInterval = time.Second * 3
)

const (
	endpoint = "https://api.open-meteo.com/v1/forecast"
)

type Sender interface {
	Send(*Data) error
}

type SMSSender struct {
	number string
}

func NewSMSSender(number string) *SMSSender {
	return &SMSSender{
		number: number,
	}
}

func (s *SMSSender) Send(data *Data) error {
	fmt.Println("sending weather to number:", s.number)
	return nil
}

type EmailSender struct {
	email string
}

func NewEmailSender(email string) *EmailSender {
	return &EmailSender{
		email: email,
	}
}

func (e *EmailSender) Send(data *Data) error {
	fmt.Println("sending weather to number:", e.email)
	return nil
}

type Data struct {
	Elevation float64        `json:"elevation"`
	Hourly    map[string]any `json:"hourly"`
}

type WPoller struct {
	senders []Sender
	closech chan struct{}
}

func NewWPoller(sender ...Sender) *WPoller {
	return &WPoller{
		senders: append(sender, sender...),
		closech: make(chan struct{}),
	}
}

func (ws *WPoller) close() {
	close(ws.closech)
}

func (ws *WPoller) start() {
	fmt.Println("starting the weather poller")
	ticker := time.NewTicker(pollInterval)
free:
	for {
		select {
		case <-ticker.C:
			data, err := GetWeatherResults(52.52, 13.41)
			if err != nil {
				log.Fatal(err)
			}
			if err := ws.handleResults(data); err != nil {
				log.Fatal(err)
			}
		case <-ws.closech:
			break free
		}
	}
}

func (wp *WPoller) handleResults(data *Data) error {
	for _, s := range wp.senders {
		if err := s.Send(data); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
func main() {
	emailSender := NewEmailSender("gg@gg.com")
	wpoller := NewWPoller(emailSender)

	go func() {
		wpoller.start()
	}()

	time.Sleep(time.Second * 3)
	wpoller.close()

	select {}

}

func GetWeatherResults(lat, long float64) (*Data, error) {
	url := fmt.Sprintf("%s?latitude=%.2f&longitude=%.2f&hourly=temperature_2m", endpoint, lat, long)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// resp, err := http.Get(endpoint)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	data := Data{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}
	return &data, nil
}
