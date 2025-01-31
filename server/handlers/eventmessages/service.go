package eventmessages

import (
	"context"
	"portal/models"
)

func (s *Service) Create(c context.Context, item *models.EventMessage) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll(c context.Context) (models.EventMessagesWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.EventMessage, error) {
	return R.Get(c, id)
}

func (s *Service) Update(c context.Context, item *models.EventMessage) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}
