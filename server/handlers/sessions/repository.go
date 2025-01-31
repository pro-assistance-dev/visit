package sessions

import (
	"context"
	"portal/models"
)

func (r *Repository) Create(c context.Context, item *models.Session) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.SessionsWithCount, err error) {
	items.Sessions = make(models.Sessions, 0)
	q := r.helper.DB.IDB(c).NewSelect().Model(&items.Sessions)
	r.helper.SQL.ExtractFTSP(c).HandleQuery(q)
	items.Count, err = q.ScanAndCount(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Session, error) {
	item := models.Session{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Session{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Session) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
