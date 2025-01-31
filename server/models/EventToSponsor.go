package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EventToSponsor struct {
	bun.BaseModel `bun:"events_to_sponsors,alias:events_to_sponsors"`
	EventID       uuid.NullUUID `bun:",pk,type:uuid" json:"eventId"`
	SponsorID     uuid.NullUUID `bun:",pk,type:uuid" json:"sponsorId"`

	Event   *Event   `bun:"rel:belongs-to,join:event_id=id"`
	Sponsor *Sponsor `bun:"rel:belongs-to,join:sponsor_id=id"`
}
