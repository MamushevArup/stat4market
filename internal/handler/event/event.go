package event

import (
	"context"
	"github.com/MamushevArup/stat4market/internal/lib/api/response"
	"github.com/MamushevArup/stat4market/internal/models"
	"github.com/go-chi/render"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type eventInsert interface {
	Insert(ctx context.Context, event models.EventRepository) error
}

// @Summary		Save event
// @Tags			api
// @Description	Save event to the storage
// @ID				event
// @Accept			json
// @Produce		json
// @Success		201
// @Param			models.EventHandler	body		models.EventHandler	true	"Event"
// @Failure		400					{object}	response.Response
// @Failure		500					{object}	response.Response
// @Router			/api/event [post]
func Save(ci eventInsert) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event models.EventHandler

		err := render.DecodeJSON(r.Body, &event)
		if err != nil {
			render.JSON(w, r, response.Error(http.StatusBadRequest, err.Error()))
			return
		}

		eventStorage := convertToStorage(event)

		err = ci.Insert(r.Context(), eventStorage)
		if err != nil {
			render.JSON(w, r, response.Error(http.StatusInternalServerError, err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func convertToStorage(event models.EventHandler) models.EventRepository {
	const layout = "2006-01-02 15:04:05"
	t, err := time.Parse(layout, event.EventTime)
	if err != nil {
		t = time.Now()
	}

	return models.EventRepository{
		EventID:   rand.New(rand.NewSource(time.Now().UnixNano())).Intn(math.MaxInt32),
		EventType: event.EventType,
		UserID:    event.UserID,
		EventTime: t,
		Payload:   event.Payload,
	}

}
