package eventdays

import (
	"context"
	"portal/models"
)

func (r *Repository) Create(c context.Context, item *models.EventDay) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.EventDaysWithCount, err error) {
	items.EventDays = make(models.EventDays, 0)
	items.Count, err = r.helper.DB.IDB(c).NewSelect().Model(&items.EventDays).
		ScanAndCount(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.EventDay, error) {
	item := models.EventDay{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.EventDay{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.EventDay) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
