package speakers

import (
	"context"
	"portal/models"
)

func (r *Repository) Create(c context.Context, item *models.Speaker) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.SpeakersWithCount, err error) {
	items.Speakers = make(models.Speakers, 0)
	q := r.helper.DB.IDB(c).NewSelect().Model(&items.Speakers).
		Relation("Human").
		Relation("Experiences")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(q)
	items.Count, err = q.ScanAndCount(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Speaker, error) {
	item := models.Speaker{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Human.Contact.Emails").
		Relation("Human.Photo").
		Relation("Human.Contact.Phones").
		Relation("Experiences").
		Where("?TableAlias.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Speaker{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Speaker) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
