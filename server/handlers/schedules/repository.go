package schedules

import (
	"context"
	"portal/models"
)

func (r *Repository) Create(c context.Context, item *models.Schedule) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.SchedulesWithCount, err error) {
	items.Schedules = make(models.Schedules, 0)
	items.Count, err = r.helper.DB.IDB(c).NewSelect().Model(&items.Schedules).
		ScanAndCount(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Schedule, error) {
	item := models.Schedule{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Schedule{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Schedule) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
