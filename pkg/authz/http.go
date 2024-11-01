package authz

import (
	"context"
	"net/http"

	"github.com/kyverno/kyverno-envoy-plugin/pkg/server"
	"github.com/kyverno/kyverno-envoy-plugin/pkg/server/handlers"
)

func NewHttpServer(addr string) server.ServerFunc {
	return func(ctx context.Context) error {
		// create mux
		mux := http.NewServeMux()
		// register health check
		mux.Handle("/livez", handlers.Healthy(handlers.True))
		// register ready check
		mux.Handle("/readyz", handlers.Ready(handlers.True))
		// create server
		s := &http.Server{
			Addr:    addr,
			Handler: mux,
		}
		// run server
		return server.RunHttp(ctx, s, "", "")
	}
}
