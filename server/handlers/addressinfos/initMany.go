package addressinfos

import (
	"context"
	"portal/models"

	"github.com/google/uuid"
	"github.com/pro-assistance-dev/sprob/helper"
)

type IService interface {
	CreateMany(models.PostAddresses) error
	UpsertMany(models.PostAddresses) error
	Upsert(*models.PostAddress) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	createMany(models.PostAddresses) error
	upsertMany(models.PostAddresses) error
	upsert(*models.PostAddress) error
	deleteMany([]uuid.UUID) error
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

func CreateService(h *helper.Helper) *Service {
	repo := NewRepository(h)
	return NewService(repo, h)
}

func NewService(repository IRepository, h *helper.Helper) *Service {
	return &Service{repository: repository, helper: h}
}

func NewRepository(h *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: h}
}
