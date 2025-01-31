package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EventToPlace struct {
	bun.BaseModel `bun:"events_to_places,alias:events_to_places"`
	EventID       uuid.NullUUID `bun:",pk,type:uuid" json:"eventId"`
	PlaceID       uuid.NullUUID `bun:",pk,type:uuid" json:"placeId"`

	Event *Event `bun:"rel:belongs-to,join:event_id=id"`
	Place *Place `bun:"rel:belongs-to,join:place_id=id"`
}

type EventsToPlaces []*EventToPlace
