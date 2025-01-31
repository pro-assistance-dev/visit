package userseventssactivities

import (
	"portal/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.UserEventActivity) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.UserEventActivity) (err error) {
	_, err = r.DB().NewInsert().On("conflict (id) do update").
		Set("end_time = current_timestamp").
		Model(item).
		Exec(r.ctx)
	return err
}

func (r *Repository) GetLastActivity(userID uuid.UUID) (*models.UserEventActivity, error) {
	item := models.UserEventActivity{}
	err := r.DB().NewSelect().Model(&item).
		Where("?TableAlias.user_id = ?", userID).
		Order("start_time desc").
		Limit(1).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) UsersActivitiesTimesSums() (items models.UsersActivitiesTimesSums, err error) {
	err = r.DB().NewSelect().Model(&items).
		Relation("User").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) UniqueUsersByHours() (items models.UniqueUsersByHours, err error) {
	err = r.DB().NewSelect().Model(&items).
		Scan(r.ctx)
	return items, err
}
