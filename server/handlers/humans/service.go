package humans

import (
	"context"
	"portal/models"
)

func (s *Service) Create(c context.Context, item *models.Human) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.Human) error {
	err := R.Update(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.HumansWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, slug string) (*models.Human, error) {
	item, err := R.Get(c, slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}
