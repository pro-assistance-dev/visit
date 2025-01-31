package customsections

import (
	"context"
	"portal/models"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.CustomSection) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.CustomSectionsWithCount, err error) {
	item.CustomSections = make(models.CustomSections, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.CustomSections).Order("item_order")

	// r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(c context.Context, slug string) (*models.CustomSection, error) {
	item := models.CustomSection{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.id = ?", slug).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.CustomSection{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(c context.Context, item *models.CustomSection) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
