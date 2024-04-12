package main

import (
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type grpcServer struct {
	addr string
}

func NewGRPCServer(addr string) *grpcServer {
	return &grpcServer{
		addr: addr,
	}
}

func (s *grpcServer) Run() error {
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		slog.Error("grpc failed to listen", "error", err)
		return err
	}
	gServer := grpc.NewServer()

	slog.Info("Starting grpc server on", "the listen address", s.addr)

	return gServer.Serve(listen)
}
