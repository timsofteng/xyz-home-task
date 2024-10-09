package httpServer

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/timsofteng/xyz-home-task/internal/logger"
)

type Server struct {
	server *http.Server
}

type Cfg struct {
	Host string
	Port string
}

func New(
	ctx context.Context,
	logger logger.Logger,
	handlers StrictServerInterface,
	cfg Cfg,
) (*Server, error) {

	mux := http.NewServeMux()
	h, err := WrapToOapiHandler(logger, mux, handlers)

	if err != nil {
		return nil, fmt.Errorf("failed to wrap handlers to oapi handler: %w", err)
	}

	server := &http.Server{
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
		Addr:         cfg.Host + ":" + cfg.Port,
		Handler:      Cors(h),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	return &Server{server: server}, nil
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
