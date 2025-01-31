package models

import (
	"time"

	"github.com/uptrace/bun"
)

type UniqueUsersByHour struct {
	bun.BaseModel `bun:",alias:users_activities_times_sums"`
	DateHour      time.Time `json:"dateHour"`
	UsersCount    uint      `json:"usersCount"`
}

type UniqueUsersByHours []*UniqueUsersByHour
