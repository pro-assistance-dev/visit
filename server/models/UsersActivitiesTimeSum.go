package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UsersActivitiesTimesSum struct {
	bun.BaseModel `bun:"users_activities_times_sums,alias:users_activities_times_sums"`
	UserID        uuid.NullUUID `json:"userId"`
	User          *User         `bun:"rel:belongs-to" json:"user"`
	TimeSum       time.Time     `json:"timeSum"`
	TimeDay       time.Time     `json:"timeDay"`
}

type UsersActivitiesTimesSums []*UsersActivitiesTimesSum
