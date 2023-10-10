package app

import (
	"fmt"
	"github.com/akhmettolegen/translator/config"
	v1 "github.com/akhmettolegen/translator/internal/controller/http/v1"
	"github.com/akhmettolegen/translator/internal/usecase"
	"github.com/akhmettolegen/translator/internal/usecase/repo"
	"github.com/akhmettolegen/translator/internal/usecase/webapi"
	"github.com/akhmettolegen/translator/pkg/httpserver"
	"github.com/akhmettolegen/translator/pkg/logger"
	"github.com/akhmettolegen/translator/pkg/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	translationUseCase := usecase.New(
		repo.New(pg),
		webapi.New(),
	)

	// HTTP Server
	router := setupRouter(l, translationUseCase)
	httpServer := httpserver.New(router, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func setupRouter(l logger.Interface, t usecase.Translation) chi.Router {
	router := chi.NewRouter()
	// Options
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// TODO: Swagger

	router.Get("/health", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })

	// TODO: Prometheus metrics

	// Routes
	router.Mount("/v1/translation", v1.NewTranslationRoutes(t, l).Routes())

	return router
}