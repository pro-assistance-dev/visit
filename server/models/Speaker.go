package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Speaker struct {
	bun.BaseModel `bun:"speakers,select:speakers_view,alias:speakers_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Position      string        `json:"position"`

	AcademicDegree   string `json:"academicDegree"`
	AcademicRank     string `json:"academicRank"`
	AcademicRelation string `json:"academicRelation"`

	Human   *Human        `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.NullUUID `bun:"type:uuid" json:"humanId"`

	//
	FullName    string      `bun:"-" json:"fullName"`
	City        string      `json:"city"`
	Experiences Experiences `bun:"rel:has-many" json:"experiences"`
}

type Speakers []*Speaker

type SpeakersWithCount struct {
	Speakers Speakers `json:"items"`
	Count    int      `json:"count"`
}
