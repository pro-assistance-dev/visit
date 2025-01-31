package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PerfomToSpeaker struct {
	bun.BaseModel `bun:"perfoms_to_speakers,alias:perfoms_to_speakers"`
	SpeakerID     uuid.NullUUID `bun:",pk,type:uuid" json:"speakerId"`
	PerfomID      uuid.NullUUID `bun:",pk,type:uuid" json:"perfomId"`

	Speaker *Speaker `bun:"rel:belongs-to,join:speaker_id=id"`
	Perfom  *Perfom  `bun:"rel:belongs-to,join:perfom_id=id"`
}

type PerfomsToSpeakers []*PerfomToSpeaker
