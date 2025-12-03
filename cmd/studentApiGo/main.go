package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/siddhesht795/studentApiGo/internal/config"
	"github.com/siddhesht795/studentApiGo/internal/http/handlers/student"
)

func main() {
	// Steps:-
	// 1. load config
	// 2. database set up
	// 3. set up router
	// 4. set up http server

	// load config
	cfg := config.MustLoad()

	// router set up
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server starting at", slog.String("Address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server", err)
		}
	}()

	<-done

	slog.Info("\nShutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown succesfuly")
}
