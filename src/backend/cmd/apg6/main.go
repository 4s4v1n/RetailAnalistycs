package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"APG6/config"
	"APG6/internal/controller/auth"
	v1 "APG6/internal/controller/http"
	"APG6/internal/repository/postgres"
	"APG6/pkg/db/postgres_db"
	"APG6/pkg/http_server"
	"APG6/pkg/logger"
)

func main() {
	logger.Init()
	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Cron("0 0 * * *").Do(logger.Init)
	s.StartAsync()

	cfg, err := config.New()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("Completed read config")

	adminDb, err := postgres_db.NewPgx(postgres_db.DSN(cfg.Postgres.DSN.Admin))
	if err != nil {
		log.Fatal().Msgf("AdminDb error: %s", err)
	}
	log.Info().Msg("Completed init adminDb")
	defer adminDb.Close()

	visitorDb, err := postgres_db.NewPgx(postgres_db.DSN(cfg.Postgres.DSN.Visitor))
	if err != nil {
		log.Fatal().Msgf("VisitorDb error: %s", err)
	}
	log.Info().Msg("Completed init visitorDb")
	defer visitorDb.Close()

	auth.Init(cfg.Postgres.Roles, cfg.Jwt)

	tableManager := postgres.New(adminDb, visitorDb)
	viewManager := postgres.New(adminDb, visitorDb)
	dataManager := postgres.New(adminDb, visitorDb)
	functionManager := postgres.New(adminDb, visitorDb)

	mux := chi.NewRouter()

	handler := v1.New(mux, tableManager, viewManager, dataManager, functionManager)
	handler.Run()

	httpServer := http_server.New(mux, http_server.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info().Msgf("Signal: %s", s.String())
	case err = <-httpServer.Notify():
		log.Error().Msgf("HttpServer.Notify: %v", err)
	}

	if err := httpServer.Shutdown(); err != nil {
		log.Error().Msgf("HttpServer.Shutdown: %v", err)
	}
}
