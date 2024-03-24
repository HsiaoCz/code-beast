package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

const (
	ttl     = time.Second * 3
	checkID = "check_health"
)

type Serveice struct {
	consulClient *api.Client
}

func NewService() *Serveice {
	client, err := api.NewClient(&api.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &Serveice{
		consulClient: client,
	}
}

func (s *Serveice) Start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	go s.registerService()
	go s.UpdateHealthCheck()
	go s.acceptLoop(ln)
}

func (s *Serveice) UpdateHealthCheck() {
	ticker := time.NewTicker(time.Second * 2)
	for {
		s.consulClient.Agent().UpdateTTL(checkID, "online", api.HealthPassing)
		<-ticker.C
	}
}

func (s *Serveice) acceptLoop(ln net.Listener) {
	for {
		_, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *Serveice) registerService() {
	check := &api.AgentServiceCheck{
		DeregisterCriticalServiceAfter: ttl.String(),
		TLSSkipVerify:                  true,
		TTL:                            ttl.String(),
		CheckID:                        checkID,
	}

	register := &api.AgentServiceRegistration{
		ID:    "login_service",
		Name:  "mycluster",
		Tags:  []string{"login"},
		Port:  3000,
		Check: check,
	}

	query := map[string]any{
		"type": "service",
	}

	plan, err := watch.Parse(query)
	if err != nil {
		log.Fatal(err)
	}

	plan.HybridHandler = func(bpv watch.BlockingParamVal, i interface{}) {
		switch msg := i.(type) {
		case []*api.ServiceEntry:
			for _, entry := range msg {
				fmt.Println("new member join", entry)
			}
		}
		fmt.Println(i)
	}
	go func() {
		plan.RunWithConfig("", &api.Config{})
	}()
	if err := s.consulClient.Agent().ServiceRegister(register); err != nil {
		log.Fatal(err)
	}
}

func main() {
	server := NewService()
	server.Start()
}
