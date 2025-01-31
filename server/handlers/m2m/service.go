package m2m

import (
	"context"
)

func (s *Service) Create(c context.Context, item *M2M) error {
	return R.Create(c, item)
}

func (s *Service) Update(c context.Context, item *M2M) error {
	return R.Update(c, item)
}

func (s *Service) Delete(c context.Context, item *M2M) error {
	return R.Delete(c, item)
}
