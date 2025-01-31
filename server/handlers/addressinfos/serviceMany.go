package addressinfos

import (
	"portal/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.PostAddresses) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items models.PostAddresses) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}

func (s *Service) Upsert(items *models.PostAddress) error {
	return s.repository.upsert(items)
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
