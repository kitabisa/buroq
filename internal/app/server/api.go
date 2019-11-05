package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/kitabisa/go-bootstrap/config"
	"github.com/kitabisa/go-bootstrap/internal/app/service"
)

// IServer interface for server
type IServer interface {
	StartApp()
	StartMetric()
}

type server struct {
	config  config.Provider
	service *service.Service
}

// NewServer create object server
func NewServer(config config.Provider, service *service.Service) IServer {
	return &server{
		config:  config,
		service: service,
	}
}

func (s *server) StartApp() {
	var srv http.Server
	idleConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// TODO: use logger
		fmt.Println("Server is shutting down")

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			// TODO: use logger
			fmt.Printf("Fail to shutting down: %v", err)
		}
		close(idleConnectionClosed)
	}()

	srv.Addr = fmt.Sprintf("%s:%d", s.config.GetString("app.host"), s.config.GetInt("app.port"))
	// srv.Handler = TODO: chi based handler

	fmt.Printf("HTTP serve at %s\n", srv.Addr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		// TODO: use logger
		fmt.Printf("Fail to start listen and server: %v", err)
	}

	<-idleConnectionClosed
	// TODO: use logger
	fmt.Println("Bye")
}

func (s *server) StartMetric() {

}
