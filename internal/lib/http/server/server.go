package server

import (
	"fmt"
	"github.com/MamushevArup/stat4market/config"
	"github.com/MamushevArup/stat4market/internal/handler"
	"net/http"
	"strconv"
	"time"
)

func New(cfg *config.Config, handler handler.Handler) (*http.Server, error) {

	if cfg.HTTP.Port <= 0 {
		cfg.HTTP.Port = 8080
	}

	port := strconv.Itoa(cfg.HTTP.Port)

	headerTimeout, err := convertConfigToDuration(cfg.HTTP.HeaderTimeout)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	idleTimeout, err := convertConfigToDuration(cfg.HTTP.IdleTimeout)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	timeout, err := convertConfigToDuration(cfg.HTTP.Timeout)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &http.Server{
		Addr:              ":" + port,
		Handler:           handler.Routes(),
		ReadTimeout:       timeout,
		ReadHeaderTimeout: headerTimeout,
		IdleTimeout:       idleTimeout,
	}, nil
}

func convertConfigToDuration(cfg string) (time.Duration, error) {
	res, err := time.ParseDuration(cfg)
	if err != nil {
		return 0, fmt.Errorf("unable to parse duration %w", err)
	}
	return res, nil
}
