package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Session struct {
	bun.BaseModel `bun:"sessions,alias:sessions"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Description   string        `json:"description"`

	Chairs     Speakers      `bun:"m2m:sessions_to_speakers,join:Session=Speaker" json:"chairs"`
	Schedule   *Schedule     `bun:"rel:belongs-to" json:"schedule"`
	ScheduleID uuid.NullUUID `bun:"type:uuid" json:"scheduleId"`

	StartTime string  `json:"startTime"`
	EndTime   string  `json:"endTime"`
	Perfoms   Perfoms `bun:"rel:has-many" json:"perfoms"`
}

type Sessions []*Session

type SessionsWithCount struct {
	Sessions Sessions `json:"items"`
	Count    int      `json:"count"`
}
