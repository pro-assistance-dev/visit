package userseventssactivities

import (
	"context"
	"portal/models"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/pro-assistance-dev/sprob/helper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Upsert(c *gin.Context)
	UsersActivitiesTimesSums(c *gin.Context)
	UniqueUsersByHours(c *gin.Context)
}

type IService interface {
	Upsert(userID uuid.UUID) error
	UsersActivitiesTimesSums() (models.UsersActivitiesTimesSums, error)
	UniqueUsersByHours() (models.UniqueUsersByHours, error)
}

type IRepository interface {
	DB() *bun.DB
	create(active *models.UserEventActivity) error
	upsert(active *models.UserEventActivity) error
	GetLastActivity(userID uuid.UUID) (*models.UserEventActivity, error)
	UsersActivitiesTimesSums() (models.UsersActivitiesTimesSums, error)
	UniqueUsersByHours() (models.UniqueUsersByHours, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.UserEventActivity, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helper.Helper
}

type Service struct {
	helper     *helper.Helper
	repository IRepository
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

type FilesService struct {
	helper *helper.Helper
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
