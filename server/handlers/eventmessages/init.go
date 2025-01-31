package eventmessages

import (
	"context"
	"portal/models"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance-dev/sprob/handlers/basehandler"
	"github.com/pro-assistance-dev/sprob/helper"
)

type IHandler interface {
	basehandler.IHandler

	FTSP(c *gin.Context)
}

type IService interface {
	basehandler.IServiceWithContext[models.EventMessage, models.EventMessages, models.EventMessagesWithCount]
}

type IRepository interface {
	basehandler.IRepositoryWithContext[models.EventMessage, models.EventMessages, models.EventMessagesWithCount]
}

type IFilesService interface {
	basehandler.IFilesService
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helper.Helper
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

type FilesService struct {
	helper *helper.Helper
}

var (
	H *Handler
	S *Service
	R *Repository
	F *FilesService
)

func Init(h *helper.Helper) {
	R = NewRepository(h)
	S = NewService(R, h)
	F = NewFilesService(h)
	H = NewHandler(S, F, h)
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helper.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
