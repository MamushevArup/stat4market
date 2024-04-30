package clickhouse

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MamushevArup/stat4market/internal/models"
	"github.com/Masterminds/squirrel"
	"time"
)

var noRow = errors.New("no rows in storage found")

func (r *Repository) Events(ctx context.Context, eventType string, from, to time.Time) ([]models.Events, error) {
	r.lg.Infof("clickhouse.Event executed")
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	var result []models.Events

	query, args, err := sq.
		Select("eventID", "eventType").
		From("events").
		Where(squirrel.And{
			squirrel.Eq{"eventType": eventType},
			squirrel.GtOrEq{"eventTime": from},
			squirrel.LtOrEq{"eventTime": to},
		}).
		ToSql()

	if err != nil {
		r.lg.Errorf("fail build query for event type=%v, err=%v", eventType, err)
		return result, fmt.Errorf("failed to build select query: %w", err)
	}

	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, noRow
		}
		return nil, fmt.Errorf("fail to get row from storage due to %w", err)
	}

	for rows.Next() {
		var eventID int
		var event string

		err = rows.Scan(&eventID, &event)
		if err != nil {
			r.lg.Errorf("fail scan eventID=%v, err=%v", eventID, err)
			return nil, fmt.Errorf("unable to match the scan %w", err)
		}

		result = append(result, models.Events{EventID: eventID, EventType: event})
	}

	defer rows.Close() // nolint: errcheck

	return result, nil
}
