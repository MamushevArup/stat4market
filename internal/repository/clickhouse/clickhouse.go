package clickhouse

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/MamushevArup/stat4market/pkg/logger"
)

type Repository struct {
	conn driver.Conn
	lg   *logger.Logger
}

func New(conn driver.Conn, lg *logger.Logger) *Repository {
	return &Repository{conn: conn, lg: lg}
}
