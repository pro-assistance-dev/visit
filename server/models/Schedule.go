package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Schedule struct {
	bun.BaseModel `bun:"schedules,alias:schedules"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Date          time.Time     `bun:"item_date" json:"scheduleDate"`

	EventDayID uuid.NullUUID `json:"eventDayId"`
	PlaceID    uuid.NullUUID `json:"placeId"`

	EventID uuid.NullUUID `bun:"type:uuid" json:"eventId"`

	// Perfoms            Perfoms            `bun:"m2m:schedules_to_perfoms,join:Schedule=Perfom" json:"perfoms"`
	SchedulesToPerfoms SchedulesToPerfoms `bun:"rel:has-many" json:"schedulesToPerfoms"`
	Sessions           Sessions           `bun:"rel:has-many" json:"sessions"`
	Perfoms            Perfoms            `bun:"rel:has-many" json:"perfoms"`

	Stream   *Stream       `bun:"rel:belongs-to" json:"stream"`
	StreamID uuid.NullUUID `bun:"type:uuid" json:"streamId"`
}

type Schedules []*Schedule

type SchedulesWithCount struct {
	Schedules Schedules `json:"items"`
	Count     int       `json:"count"`
}
