package models

import "github.com/google/uuid"

type Stream struct {
	ID string `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	// StatusPublish string `json:"statusPublish"`
	// StatusView    string `json:"statusView"`
	// Poster   *FileInfo     `bun:"rel:belongs-to" json:"poster"`
	PosterID uuid.NullUUID `bun:"type:uuid" json:"posterId"`
}

type Streams []Stream
