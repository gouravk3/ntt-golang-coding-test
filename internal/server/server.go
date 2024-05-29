package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gouravk3/ntt-golang-coding-test/config"
	"github.com/gouravk3/ntt-golang-coding-test/internal/router"
)

type server struct {
	config     *config.Config
	httpServer *http.Server
}

func New(config *config.Config) *server {
	ginEngine, err := router.InitRouter(config)
	if err != nil {
		panic(err)
	}

	// decoupling handler
	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port),
		Handler: ginEngine.Handler(),
	}

	return &server{
		config:     config,
		httpServer: srv,
	}
}

func (s *server) run() {
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func (s *server) gracefullShutdown() {
	defer func() {
		s.httpServer.Shutdown(context.Background())
	}()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
}

func (s *server) Start() {
	go s.run()
	s.gracefullShutdown()
}
