package repository

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	clickhousedb "github.com/MamushevArup/stat4market/internal/repository/clickhouse"
	"github.com/MamushevArup/stat4market/pkg/logger"
)

type Repository struct {
	Clickhouse *clickhousedb.Repository
}

func New(conn driver.Conn, lg *logger.Logger) *Repository {
	return &Repository{
		Clickhouse: clickhousedb.New(conn, lg),
	}
}
