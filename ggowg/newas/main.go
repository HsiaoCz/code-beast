package main

import "log/slog"

func main() {
	server := NewAPIServer(":8008")
	if err := server.Run(); err != nil {
		slog.Error("api server start error", "err", err)
		return
	}
}
