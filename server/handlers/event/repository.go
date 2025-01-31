package event

import (
	"context"
	"portal/models"

	"github.com/uptrace/bun"
)

func (r *Repository) Create(c context.Context, item *models.Event) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) CreateMessage(c context.Context, item *ChatMessage[WithId]) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.EventsWithCount, err error) {
	items.Events = make(models.Events, 0)
	q := r.helper.DB.IDB(c).NewSelect().Model(&items.Events).
		Relation("EventDays").
		Where("archive = false")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(q)
	items.Count, err = q.ScanAndCount(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Event, error) {
	item := models.Event{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Chat.User.Human").
		Relation("Perfoms").
		Relation("Places").
		Relation("EventDays").
		Relation("Schedules").
		Relation("Schedules.Sessions", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("sessions.start_time")
		}).
		Relation("Schedules.Perfoms", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				Order("start_time")
		}).
		Relation("Schedules.Perfoms.Speakers.Human").
		Relation("Schedules.Sessions.Chairs.Human").
		Relation("Sponsors.Banners.FileInfo").
		Where("events.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Event{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Event) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) GetBySlug(c context.Context, slug string) (*models.Event, error) {
	item := models.Event{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Chat.ChatMessages.User").
		Relation("Perfoms").
		Relation("Places").
		Relation("EventDays").
		Relation("Schedules").
		// Relation("EventDays.Schedules.SchedulesToPerfoms.Perfom.Speakers.Human").
		// Relation("EventDays.Schedules.Perfoms.Speakers.Human").
		Relation("Schedules.Sessions", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				Order("start_time")
		}).
		Relation("Schedules.Perfoms", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				Order("start_time")
		}).
		Relation("Schedules.Perfoms.Speakers.Human").
		Relation("Sponsors.Banners.FileInfo").
		Where("events.slug = ?", slug).
		Scan(c)
	return &item, err
}

// func (r *Repository) CreateMessage(c context.Context, item *models.EventMessage) (err error) {
// 	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
// 	return err
// }

func (r *Repository) GetAllMessages(c context.Context, eventID string) (models.EventMessages, error) {
	items := make(models.EventMessages, 0)
	lastMessagesQuery := r.helper.DB.IDB(c).NewSelect().Model((*models.EventMessage)(nil)).Relation("User").
		Where("event_messages.event_id = ?", eventID).
		Order("created_on DESC").
		Limit(30)
	err := r.helper.DB.IDB(c).NewSelect().
		ModelTableExpr("(?) AS last_messages", lastMessagesQuery).
		Order("last_messages.created_on ASC").
		Scan(c, &items)
	return items, err
}
