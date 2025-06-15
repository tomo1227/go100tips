package main

import "log/slog"

func main() {
	logger := NewLogger()
	logger.Info("Hello from logger!", slog.String("component", "main"))
}
