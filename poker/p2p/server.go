package p2p

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"
)

// type TCPTransport struct{
//}

type Peer struct {
	conn net.Conn
}

func (p *Peer) Send(b []byte) error {
	_, err := p.conn.Write(b)
	return err
}

type ServerConfig struct {
	ListenAddr string
}

type Message struct {
	Payload io.Reader
	From    net.Addr
}

type Server struct {
	ServerConfig

	listener net.Listener
	mu       sync.Mutex
	peers    map[net.Addr]*Peer
	addPeer  chan *Peer

	delPeer chan *Peer
	handler Handler
	msgch   chan *Message
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		handler:      &DefaultHandler{},
		ServerConfig: cfg,
		peers:        make(map[net.Addr]*Peer),
		addPeer:      make(chan *Peer),
		msgch:        make(chan *Message),
	}
}

func (s *Server) Start() {
	go s.loop()

	if err := s.listen(); err != nil {
		panic(err)
	}
	go s.acceptLoop()
}

func (s *Server) handleConn(p *Peer) {

	defer func() {
		s.delPeer <- p
	}()

	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			break
		}
		s.msgch <- &Message{
			From:    p.conn.RemoteAddr(),
			Payload: bytes.NewBuffer(buf[:n]),
		}
		fmt.Println(string(buf[:n]))
	}
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(err)
		}
		peer := &Peer{
			conn: conn,
		}
		s.addPeer <- peer
		peer.Send([]byte("POKER v0.0.1"))
		go s.handleConn(peer)
	}
}

func (s *Server) listen() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.listener = ln
	return nil
}

func (s *Server) loop() {
	for {
		select {
		case peer := <-s.addPeer:
			s.peers[peer.conn.RemoteAddr()] = peer
			fmt.Printf("new player connected %s\n", peer.conn.RemoteAddr())
		default:
			fmt.Println("well done")
		}
	}
}
