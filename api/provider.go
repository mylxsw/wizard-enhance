package api

import (
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
	"github.com/mylxsw/asteria/log"
	"github.com/mylxsw/glacier/infra"
	"github.com/mylxsw/glacier/listener"
	"github.com/mylxsw/glacier/web"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/mylxsw/wizard-enhance/config"
)

type Provider struct{}

func (s Provider) Aggregates() []infra.Provider {
	return []infra.Provider{
		web.Provider(
			listener.FlagContext("listen"),
			web.SetRouteHandlerOption(s.routes),
			web.SetMuxRouteHandlerOption(s.muxRoutes),
			web.SetExceptionHandlerOption(s.exceptionHandler),
		),
	}
}

func (s Provider) Register(app infra.Binder) {}
func (s Provider) Boot(app infra.Resolver)   {}

func (s Provider) exceptionHandler(ctx web.Context, err interface{}) web.Response {
	log.Errorf("error: %v, call stack: %s", err, debug.Stack())
	return nil
}

func (s Provider) muxRoutes(cc infra.Resolver, router *mux.Router) {
	cc.MustResolve(func(conf *config.Config) {
		// prometheus metrics
		router.PathPrefix("/metrics").Handler(promhttp.Handler())
		// health check
		router.PathPrefix("/health").Handler(HealthCheck{})
	})
}

func (s Provider) routes(cc infra.Resolver, router web.Router, mw web.RequestMiddleware) {
	mws := make([]web.HandlerDecorator, 0)
	mws = append(mws, mw.AccessLog(log.Module("api")), mw.CORS("*"))

	router.WithMiddleware(mws...).Controllers(
		"/api",
		controllers(cc)...,
	)
}

type HealthCheck struct{}

func (h HealthCheck) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte(`{"status": "UP"}`))
}
