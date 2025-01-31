package banners

import (
	"context"
	"portal/models"
)

func (r *Repository) Create(c context.Context, item *models.Banner) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.BannersWithCount, err error) {
	items.Banners = make(models.Banners, 0)
	q := r.helper.DB.IDB(c).NewSelect().Model(&items.Banners)
	r.helper.SQL.ExtractFTSP(c).HandleQuery(q)
	items.Count, err = q.ScanAndCount(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Banner, error) {
	item := models.Banner{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Speakers.Human").
		Where("?TableAlias.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Banner{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Banner) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
