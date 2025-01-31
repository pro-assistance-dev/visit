package userseventssactivities

import (
	"portal/models"
	"time"

	"github.com/google/uuid"
)

func (s *Service) Upsert(userID uuid.UUID) error {
	lastActivity, _ := s.repository.GetLastActivity(userID)
	if lastActivity != nil {
		now := time.Now()
		diff := now.Sub(lastActivity.End)
		if diff > time.Minute {
			userActivity := &models.UserEventActivity{UserID: uuid.NullUUID{UUID: userID, Valid: true}}
			return s.repository.create(userActivity)
		}
	}
	return s.repository.upsert(lastActivity)
}

func (s *Service) UsersActivitiesTimesSums() (models.UsersActivitiesTimesSums, error) {
	return s.repository.UsersActivitiesTimesSums()
}

func (s *Service) UniqueUsersByHours() (models.UniqueUsersByHours, error) {
	return s.repository.UniqueUsersByHours()
}
