package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/kitabisa/buroq/internal/app/commons"
	"github.com/kitabisa/buroq/internal/app/graphql"
	"github.com/kitabisa/buroq/internal/app/handler"
	"github.com/kitabisa/buroq/internal/app/service"
	"github.com/sirupsen/logrus"
)

// IServer interface for server
type IServer interface {
	StartApp()
}

type server struct {
	opt      commons.Options
	services *service.Services
}

// NewServer create object server
func NewServer(opt commons.Options, services *service.Services) IServer {
	return &server{
		opt:      opt,
		services: services,
	}
}

func (s *server) StartApp() {
	var srv http.Server
	idleConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		logrus.Infoln("[API] Server is shutting down")

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logrus.Infof("[API] Fail to shutting down: %v", err)
		}
		close(idleConnectionClosed)
	}()

	srv.Addr = fmt.Sprintf("%s:%d", s.opt.Config.GetString("app.host"), s.opt.Config.GetInt("app.port"))
	hOpt := handler.HandlerOption{
		Options:  s.opt,
		Services: s.services,
	}

	if s.opt.Config.GetBool("graphql.is_enabled") {
		logrus.Infoln("[API] GraphQL schema is enabled")
		logrus.Infoln(fmt.Sprintf("%s%s", "[API] GraphQL route: /", s.opt.Config.GetString("graphql.route")))
		schema := graphql.InitGraphqlSchema(s.services)
		hOpt.GraphqlSchema = schema
	}

	srv.Handler = Router(hOpt)

	logrus.Infof("[API] HTTP serve at %s\n", srv.Addr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		logrus.Infof("[API] Fail to start listen and server: %v", err)
	}

	<-idleConnectionClosed
	logrus.Infoln("[API] Bye")
}
