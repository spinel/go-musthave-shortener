package router

import (
	"net/http/pprof"
	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/handler"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
)

// Router for an app.
func NewRouter(cfg config.Config, repo repository.URLStorer) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.NewCreateURLHandler(cfg, repo))
	r.HandleFunc("/api/shorten", handler.NewCreateJSONURLHandler(cfg, repo))
	r.HandleFunc("/ping", handler.NewPingHandler(repo))
	r.HandleFunc("/user/urls", handler.NewGetUserURLsHandler(cfg, repo))
	r.HandleFunc("/api/shorten/batch", handler.NewCreateBatchHandler(cfg, repo))
	r.HandleFunc("/api/user/urls", handler.NewDeleteBatchHandler(cfg, repo))
	r.HandleFunc("/{id:[0-9a-z]+}", handler.NewGetURLHandler(cfg, repo))

	// pprof
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/debug/pprof/block", pprof.Handler("block"))

	return r
}
