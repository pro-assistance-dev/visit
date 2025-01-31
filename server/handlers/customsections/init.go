package customsections

import (
	"context"
	"portal/models"

	"github.com/pro-assistance-dev/sprob/handlers/basehandler"
	"github.com/pro-assistance-dev/sprob/helper"
	"github.com/uptrace/bun"
)

type IHandler interface {
	basehandler.IHandler
}

type IService interface {
	basehandler.IServiceWithContext[models.CustomSection, models.CustomSections, models.CustomSectionsWithCount]
}

type IRepository interface {
	basehandler.IRepositoryWithContext[models.CustomSection, models.CustomSections, models.CustomSectionsWithCount]
}

type IFilesService interface {
	basehandler.IFilesService
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helper.Helper
}

var (
	H *Handler
	S *Service
	R *Repository
	F *FilesService
)

type Service struct {
	Repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
	Tx     *bun.Tx
}

type FilesService struct {
	helper *helper.Helper
}

func Init(h *helper.Helper) {
	R = NewRepository(h)
	S = NewService(R, h)
	F = NewFilesService(h)
	H = NewHandler(S, F, h)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helper.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{Repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
