package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EventDay struct {
	bun.BaseModel `bun:"event_days,alias:event_days"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Date          time.Time     `bun:"item_date" json:"date"`
	Description   string        `json:"description"`

	EventID uuid.NullUUID `json:"eventId"`
}

type EventDays []*EventDay

type EventDaysWithCount struct {
	EventDays EventDays `json:"items"`
	Count     int       `json:"count"`
}
