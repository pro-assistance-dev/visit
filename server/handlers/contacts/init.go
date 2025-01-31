package contacts

import (
	"portal/models"

	"github.com/pro-assistance-dev/sprob/handlers/basehandler"
	"github.com/pro-assistance-dev/sprob/helper"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	FTSP(c *gin.Context)
	basehandler.IHandler
}

type IService interface {
	basehandler.IServiceWithContext[models.Contact, models.Contacts, models.ContactsWithCount]
}

type IRepository interface {
	basehandler.IRepositoryWithContext[models.Contact, models.Contacts, models.ContactsWithCount]
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
