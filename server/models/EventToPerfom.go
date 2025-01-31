package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EventToPerfom struct {
	bun.BaseModel `bun:"events_to_perfoms,alias:events_to_perfoms"`
	EventID       uuid.NullUUID `bun:",pk,type:uuid" json:"eventId"`
	PerfomID      uuid.NullUUID `bun:",pk,type:uuid" json:"perfomId"`

	Event  *Event  `bun:"rel:belongs-to,join:event_id=id"`
	Perfom *Perfom `bun:"rel:belongs-to,join:perfom_id=id" json:"perform"`
}

type EventsToPerfoms []*EventToPerfom
