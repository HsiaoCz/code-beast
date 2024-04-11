package main

import (
	"fmt"
	"log/slog"
	"net"
)

const defaultAddr = ":5002"

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	peers     map[*Peer]bool
	ln        net.Listener
	addPeerch chan *Peer
	quitch    chan struct{}
}

func NewServer(cfg Config) *Server {
	return &Server{
		Config:    cfg,
		peers:     make(map[*Peer]bool),
		addPeerch: make(chan *Peer),
		quitch:    make(chan struct{}),
	}
}

func (s *Server) Start() error {
	if len(s.ListenAddr) == 0 {
		s.ListenAddr = defaultAddr
	}
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		slog.Error("the server listen error", "err", err)
		return err
	}
	s.ln = ln

	go s.loop()

	slog.Info("server running", "listenAddr", s.ListenAddr)

	return s.acceptLoop()
}

func (s *Server) loop() {
	for {
		select {
		case <-s.quitch:
			return
		case peer := <-s.addPeerch:
			s.peers[peer] = true
		default:
			fmt.Println("foo")
		}
	}
}

func (s *Server) acceptLoop() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("accept error", "error", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	peer := NewPeer(conn)
	s.addPeerch <- peer
	slog.Info("new peer connected", "remote addr", conn.RemoteAddr())
	if err := peer.readLoop(); err != nil {
		slog.Error("peer read remote error", "err", err, "remote addr", conn.RemoteAddr())
		return
	}
}

func main() {

}
