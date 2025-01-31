package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EventMessage struct {
	bun.BaseModel `bun:"event_messages,alias:event_messages"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Message       string        `json:"message"`
	CreatedOn     time.Time     `bun:",nullzero,notnull" json:"createdOn"`
	UserID        *string       `json:"userId"`
	User          *User         `bun:"rel:belongs-to" json:"user"`
	EventID       *string       `json:"eventId"`
}

type EventMessages []EventMessage
type EventMessagesWithCount struct {
	EventMessages EventMessages `json:"items"`
	Count         int           `json:"count"`
}
