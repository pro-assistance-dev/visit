package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserEventActivity struct {
	bun.BaseModel `bun:"users_events_activities,alias:users_events_activities"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	UserID        uuid.NullUUID `json:"userId"`
	User          *User         `bun:"rel:belongs-to" json:"user"`
	Start         time.Time     `bun:"start_time,nullzero,notnull,default:current_timestamp"`
	End           time.Time     `bun:"end_time,nullzero,notnull,default:current_timestamp"`
}
