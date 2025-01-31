package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Perfom struct {
	bun.BaseModel `bun:"perfoms,alias:perfoms"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Description   string        `json:"description"`

	SchedulesToPerfoms SchedulesToPerfoms `bun:"rel:has-many" json:"schedulesToPerfoms"`
	Speakers           Speakers           `bun:"m2m:perfoms_to_speakers,join:Perfom=Speaker" json:"speakers"`

	StartTime  string        `json:"startTime"`
	EndTime    string        `json:"endTime"`
	SessionID  uuid.NullUUID `json:"sessionId"`
	EventID    uuid.NullUUID `json:"eventId"`
	ScheduleID uuid.NullUUID `json:"scheduleId"`
	Type       string        `json:"type"`
	Format     string        `json:"format"`
}

type Perfoms []*Perfom

type PerfomsWithCount struct {
	Perfoms Perfoms `json:"items"`
	Count   int     `json:"count"`
}
