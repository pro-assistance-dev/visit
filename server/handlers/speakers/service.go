package speakers

import (
	"context"
	"portal/models"
)

func (s *Service) Create(c context.Context, item *models.Speaker) error {
	return R.Create(c, item)
}

func (s *Service) GetAll(c context.Context) (models.SpeakersWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.Speaker, error) {
	return R.Get(c, id)
}

func (s *Service) Update(c context.Context, item *models.Speaker) error {
	return R.Update(c, item)
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}
