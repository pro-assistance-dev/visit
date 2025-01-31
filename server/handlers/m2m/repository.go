package m2m

import (
	"context"

	"github.com/uptrace/bun"
)

func createModelMap(item *M2M) map[string]interface{} {
	model := map[string]interface{}{
		item.ID1.Col: item.ID1.Value,
		item.ID2.Col: item.ID2.Value,
	}

	for f := range item.Fields {
		model[item.Fields[f].Col] = item.Fields[f].Value
	}
	return model
}

func (r *Repository) Create(c context.Context, item *M2M) (err error) {
	model := createModelMap(item)
	_, err = r.helper.DB.IDB(c).NewInsert().Model(&model).TableExpr(item.TableName).
		Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *M2M) (err error) {
	model := createModelMap(item)
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(&model).TableExpr(item.TableName).
		Where("? = ?", bun.Ident(item.ID1.Col), item.ID1.Value).
		Where("? = ?", bun.Ident(item.ID2.Col), item.ID2.Value).
		Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, item *M2M) (err error) {
	model := createModelMap(item)
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&model).TableExpr(item.TableName).
		Where("? = ?", bun.Ident(item.ID1.Col), item.ID1.Value).
		Where("? = ?", bun.Ident(item.ID2.Col), item.ID2.Value).
		Exec(c)
	return err
}
