package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Place struct {
	bun.BaseModel `bun:"places,alias:places"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
}

type Places []*Place

type PlacesWithCount struct {
	Places Places `json:"items"`
	Count  int    `json:"count"`
}
