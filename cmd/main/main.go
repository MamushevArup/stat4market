package main

import (
	"context"
	"github.com/MamushevArup/stat4market/cmd/migration"
	"github.com/MamushevArup/stat4market/config"
	"github.com/MamushevArup/stat4market/internal/handler"
	"github.com/MamushevArup/stat4market/internal/lib/http/server"
	"github.com/MamushevArup/stat4market/internal/repository"
	clickhousedb "github.com/MamushevArup/stat4market/pkg/clickhouse"
	"github.com/MamushevArup/stat4market/pkg/logger"
	"github.com/joho/godotenv"
	lg "log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const configPath = "config/config.yml"
const contextTimeout = 5 * time.Second

//	@title			Stat4Market API integration with clickhouse
//	@version		1.0
//	@description	In this test task one endpoint and three clickhouse query appear.

// @host	localhost:4444
// @scheme	http
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer func() {
		cancel()
	}()

	if err := godotenv.Load(); err != nil {
		lg.Fatal("Error loading .env file " + err.Error())
	}

	log := logger.New()

	err := migration.Run()
	if err != nil {
		lg.Fatalf("migration fail %v,", err)
		return
	}
	log.Info("migration done")

	cfg, err := config.New(configPath)
	if err != nil {
		log.Fatalf("fail to load config %v", err)
	}
	log.Info("config loaded successfully")

	conn, err := clickhousedb.New(ctx, cfg, log)
	if err != nil {
		log.Fatalf("clickhouse connection error: %v", err)
	}

	log.Infof("clickhouse connection established run on port=%v", cfg.ClickHouse.Port)
	defer func() { _ = conn.Close() }() // nolint: errcheck

	r := repository.New(conn, log)

	err = r.Clickhouse.InsertTestData(ctx)
	if err != nil {
		log.Errorf("failed to insert test data: %v", err)
	}

	log.Info("test data inserted")

	hdl := handler.New(r)

	srv, err := server.New(cfg, *hdl)
	if err != nil {
		log.Fatalf("fail to create server %v", err)
	}

	log.Infof("server started on port=%v", cfg.HTTP.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("fail to start server %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Waiting for SIGINT (pkill -2)
	<-stop

	if err = srv.Shutdown(ctx); err != nil {
		log.Errorf("unable to shutdown server %v", err)
	}

	log.Info("Server stopped")
}
