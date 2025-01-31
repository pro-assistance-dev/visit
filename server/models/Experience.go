package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Experience struct {
	bun.BaseModel `bun:"experiences,alias:experiences"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	SpeakerID uuid.NullUUID `bun:"type:uuid" json:"speakerId"`

	Start    time.Time `bun:"item_start" json:"start"`
	End      time.Time `bun:"item_end" json:"end"`
	Place    string    `json:"place"`
	Position string    `json:"position"`
	Division string    `json:"division"`
}

type Experiences []*Experience

type ExperiencesWithCount struct {
	Experiences Experiences `json:"items"`
	Count       int         `json:"count"`
}
