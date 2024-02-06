package http

import (
	"context"
	"github.com/hell-kitchen/api-gateway/internal/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Server is HTTP server struct. Includes logger, router, service, etc.
type Server struct {
	router *echo.Echo
	config *config.Server
	log    *zap.Logger
}

// NewServer return new HTTP Server.
func NewServer(config *config.Server, logger *zap.Logger) (*Server, error) {
	r := echo.New()
	r.HideBanner = true

	s := &Server{
		router: r,
		config: config,
		log:    logger,
	}

	return s, nil
}

func (srv *Server) OnStart(_ context.Context) error {
	srv.log.Debug("OnStart hook called")
	go func() {
		err := srv.router.Start(srv.config.Addr())
		if err != nil {
			srv.log.Error("unknown error while starting http server", zap.Error(err))
		}
	}()
	return nil
}

func (srv *Server) OnStop(ctx context.Context) error {
	srv.log.Debug("OnStop hook called")
	err := srv.router.Shutdown(ctx)
	srv.log.Info("stopped HTTP server", zap.Error(err))
	return err
}
