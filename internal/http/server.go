package httpserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	logger     *slog.Logger
}

func New(port string, handler http.Handler, logger *slog.Logger) *Server {
	return &Server{
		httpServer: &http.Server{Addr: port, Handler: handler, ReadHeaderTimeout: 5 * time.Second},
		logger:     logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	errch := make(chan error, 1)

	go func() {
		s.logger.Info("Starting server ...")
		if err := s.httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			errch <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		s.logger.Info("Stopping server...")
		if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return nil
	case err := <-errch:
		return err
	}
}
