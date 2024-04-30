package clickhouse

import (
	"context"
	"fmt"
	"github.com/MamushevArup/stat4market/internal/models"
	"github.com/Masterminds/squirrel"
	"math/rand"
	"time"
)

var eventTypes = []string{
	"PageLoad",
	"ButtonClick",
	"FormSubmit",
	"UserLogin",
	"ErrorOccurred",
}

func (r *Repository) InsertTestData(ctx context.Context) error {
	clickdb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	query := clickdb.
		Insert("events").
		Columns("eventID", "eventType", "userID", "eventTime", "payload")

	for i := 1; i <= 15; i++ {
		eventID := rand.Intn(1000)
		eventType := eventTypes[rand.Intn(len(eventTypes))]
		userID := rand.Intn(1000)
		eventTime := time.Now()
		payload := fmt.Sprintf("payload%d", i)

		query = query.Values(eventID, eventType, userID, eventTime, payload)
	}

	sqlQuery, args, err := query.ToSql()

	if err != nil {
		r.lg.Errorf("failed to build insert query: %v", err)
		return fmt.Errorf("failed to build insert query: %w", err)
	}

	err = r.conn.Exec(ctx, sqlQuery, args...)
	if err != nil {
		return fmt.Errorf("failed to insert test data: %w", err)
	}

	r.lg.Info("Test data inserted into events table")
	return nil
}

func (r *Repository) Insert(ctx context.Context, event models.EventRepository) error {

	clickdb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := clickdb.
		Insert("events").
		Columns("eventID", "eventType", "userID", "eventTime", "payload").
		Values(event.EventID, event.EventType, event.UserID, event.EventTime, event.Payload).
		ToSql()

	if err != nil {
		r.lg.Errorf("failed to build insert query eventID=%v, userID=%v: %v", event.EventID, event.UserID, err)
		return fmt.Errorf("failed to build insert query: %w", err)
	}

	err = r.conn.Exec(ctx, query, args...)
	if err != nil {
		r.lg.Errorf("failed to insert eventID=%v, userID=%v: %v", event.EventID, event.UserID, err)
		return fmt.Errorf("failed to insert : %w", err)
	}

	return nil
}
