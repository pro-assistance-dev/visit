package models

import (
	"github.com/google/uuid"
	baseModels "github.com/pro-assistance-dev/sprob/models"
	"github.com/uptrace/bun"
)

type Banner struct {
	bun.BaseModel `bun:"banners,alias:banners"`
	ID            uuid.NullUUID        `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FileInfo      *baseModels.FileInfo `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID    uuid.NullUUID        `bun:"type:uuid" json:"fileInfoId"`
	Sponsor       *Sponsor             `bun:"rel:belongs-to" json:"sponsor"`
	SponsorID     uuid.NullUUID        `bun:"type:uuid" json:"sponsorId"`
}

type Banners []*Banner

type BannersWithCount struct {
	Banners Banners `json:"items"`
	Count   int     `json:"count"`
}
