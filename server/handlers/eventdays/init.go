package eventdays

import (
	"portal/models"

	"github.com/pro-assistance-dev/sprob/handlers/basehandler"
	"github.com/pro-assistance-dev/sprob/helper"
)

type IHandler interface {
	basehandler.IHandler
}

type IService interface {
	basehandler.IServiceWithContext[models.EventDay, models.EventDays, models.EventDaysWithCount]
}

type IRepository interface {
	basehandler.IRepositoryWithContext[models.EventDay, models.EventDays, models.EventDaysWithCount]
}

type Handler struct {
	helper *helper.Helper
}

type Service struct {
	helper *helper.Helper
}

type Repository struct {
	helper *helper.Helper
}

var (
	H *Handler
	S *Service
	R *Repository
)

func Init(h *helper.Helper) {
	H = &Handler{helper: h}
	S = &Service{helper: h}
	R = &Repository{helper: h}
}
