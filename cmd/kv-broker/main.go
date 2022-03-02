package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/andreykont/KVStorage/api"
	"github.com/andreykont/KVStorage/broker"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const ListenAddr = "127.0.0.1:8081"

func main() {
	r := chi.NewRouter()
	mybroker := broker.NewBroker()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.SetHeader("Accept", "application/json"))
	r.Route("/api", func(r chi.Router) {
		r.Mount("/", api.Handler(mybroker))
	})

	server := &http.Server{
		Addr:    ListenAddr,
		Handler: r,
	}
	go func() {
		serverErr := server.ListenAndServe()
		if serverErr != nil {
			fmt.Printf("Can't start http server: %w", serverErr)
		}
	}()
	<-ctx.Done()
	stop()
	fmt.Printf("shutting down gracefully")
}
