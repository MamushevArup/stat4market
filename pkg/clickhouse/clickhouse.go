package clickhouse

import (
	"context"
	"errors"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/MamushevArup/stat4market/config"
	"github.com/MamushevArup/stat4market/pkg/logger"
)

func New(ctx context.Context, cfg *config.Config, log *logger.Logger) (driver.Conn, error) {
	addr := fmt.Sprintf("%s:%d", cfg.ClickHouse.Host, cfg.ClickHouse.Port)

	var (
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{addr},
			Auth: clickhouse.Auth{
				Database: cfg.ClickHouse.Database,
				Password: cfg.ClickHouse.Password,
				Username: cfg.ClickHouse.User,
			},
			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
		})
	)

	if err != nil {
		return nil, fmt.Errorf("clickhouse open error: %w", err)
	}

	if err = conn.Ping(ctx); err != nil {
		var exception *clickhouse.Exception
		if errors.As(err, &exception) {
			log.Errorf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		log.Errorf("clickhouse ping error at port=%v", cfg.ClickHouse.Port)
		return nil, fmt.Errorf("clickhouse ping error: %w", err)
	}
	return conn, nil
}
