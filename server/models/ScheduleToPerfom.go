package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ScheduleToPerfom struct {
	bun.BaseModel `bun:"schedules_to_perfoms,alias:schedules_to_perfoms"`
	ScheduleID    uuid.NullUUID `bun:",pk,type:uuid" json:"scheduleId"`
	PerfomID      uuid.NullUUID `bun:",pk,type:uuid" json:"perfomId"`

	Schedule *Schedule `bun:"rel:belongs-to,join:schedule_id=id"`
	Perfom   *Perfom   `bun:"rel:belongs-to,join:perfom_id=id" json:"perfom"`

	StartTime string        `bun:"-" json:"startTime"`
	EndTime   string        `bun:"-" json:"endTime"`
	SessionID uuid.NullUUID `bun:"-" json:"sessionId"`
}

type SchedulesToPerfoms []*ScheduleToPerfom
