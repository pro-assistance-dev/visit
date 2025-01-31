package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SessionToSpeaker struct {
	bun.BaseModel `bun:"sessions_to_speakers,alias:sessions_to_speakers"`
	SessionID     uuid.NullUUID `bun:",pk,type:uuid" json:"sessionId"`
	SpeakerID     uuid.NullUUID `bun:",pk,type:uuid" json:"speakerId"`

	Speaker *Speaker `bun:"rel:belongs-to,join:speaker_id=id"`
	Session *Session `bun:"rel:belongs-to,join:session_id=id" json:"session"`
}

type SessionsToSpeakers []*SessionToSpeaker
