package sponsors

import (
	"context"
	"portal/models"
)

func (r *Repository) Create(c context.Context, item *models.Sponsor) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.SponsorsWithCount, err error) {
	items.Sponsors = make(models.Sponsors, 0)
	q := r.helper.DB.IDB(c).NewSelect().Model(&items.Sponsors)
	r.helper.SQL.ExtractFTSP(c).HandleQuery(q)
	items.Count, err = q.ScanAndCount(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Sponsor, error) {
	item := models.Sponsor{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Banners.FileInfo").
		Where("?TableAlias.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Sponsor{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Sponsor) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
