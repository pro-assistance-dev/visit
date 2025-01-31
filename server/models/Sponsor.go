package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Sponsor struct {
	bun.BaseModel `bun:"sponsors,alias:sponsors"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	Address       string        `json:"address"`
	Description   string        `json:"description"`
	Site          string        `json:"site"`

	Banners Banners `bun:"rel:has-many" json:"banners"`
}

type Sponsors []*Sponsor

type SponsorsWithCount struct {
	Sponsors Sponsors `json:"items"`
	Count    int      `json:"count"`
}
