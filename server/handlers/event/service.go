package event

import (
	"context"
	"portal/models"
)

func (s *Service) Create(c context.Context, item *models.Event) error {
	return R.Create(c, item)
}

func (s *Service) GetAll(c context.Context) (models.EventsWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.Event, error) {
	return R.Get(c, id)
}

func (s *Service) Update(c context.Context, item *models.Event) error {
	return R.Update(c, item)
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) GetBySlug(c context.Context, id string) (*models.Event, error) {
	return R.GetBySlug(c, id)
}

func (s *Service) CreateMessage(c context.Context, item *ChatMessage[WithId]) error {
	return R.CreateMessage(c, item)
}

func (s *Service) GetAllMessages(c context.Context, eventID string) (models.EventMessages, error) {
	return R.GetAllMessages(c, eventID)
}
