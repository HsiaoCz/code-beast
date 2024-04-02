package network

import "fmt"

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts

	rpcCh  chan RPC
	quitch chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
	}
}

func (s *Server) Start() {
	s.initTransport()
free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v", rpc)
		case <-s.quitch:
			break free
		}
	}
	fmt.Println("Server shutdown")
}

func (s *Server) initTransport() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
