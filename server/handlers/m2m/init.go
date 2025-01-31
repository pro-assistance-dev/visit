package m2m

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance-dev/sprob/helper"
)

type M2M struct {
	Model     string   `json:"model"`
	ID1       keyValue `json:"id1"`
	ID2       keyValue `json:"id2"`
	TableName string
	Fields    []*keyValue `json:"fields"`
}

type keyValue struct {
	Field string `json:"field"`
	Col   string
	Value string `json:"value"`
}

type IHandler interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type IService interface {
	Create(context.Context, *M2M) error
	Update(context.Context, *M2M) error
	Delete(context.Context, *M2M) error
}

type IRepository interface {
	Create(context.Context, *M2M) error
	Update(context.Context, *M2M) error
	Delete(context.Context, *M2M) error
}

type Handler struct {
	service IService
	helper  *helper.Helper
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

var (
	H *Handler
	S *Service
	R *Repository
)

func Init(h *helper.Helper) {
	R = NewRepository(h)
	S = NewService(R, h)
	H = NewHandler(S, h)
}

func CreateHandler(h *helper.Helper) *Handler {
	repo := NewRepository(h)
	service := NewService(repo, h)
	return NewHandler(service, h)
}

// NewHandler constructor
func NewHandler(service IService, h *helper.Helper) *Handler {
	return &Handler{service: service, helper: h}
}

func NewService(repository IRepository, h *helper.Helper) *Service {
	return &Service{repository: repository, helper: h}
}

func NewRepository(h *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: h}
}
