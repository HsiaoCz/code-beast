package main

import "log/slog"

func main() {
	grpcsrv := NewGRPCServer(":9001")
	if err := grpcsrv.Run(); err != nil {
		slog.Error("grpc run error", "error", err)
		return
	}
}
