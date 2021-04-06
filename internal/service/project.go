package service

import (
	"context"
	"database/sql"

	"github.com/mylxsw/eloquent/query"
	"github.com/mylxsw/glacier/infra"
	"github.com/mylxsw/wizard-enhance/internal/service/model"
)

type ProjectService interface {
	GetByID(ctx context.Context, id int64) (model.Project, error)
}

func NewProjectService(cc infra.Resolver, db *sql.DB) ProjectService {
	return &projectService{db: db, cc: cc}
}

type projectService struct {
	db *sql.DB
	cc infra.Resolver
}

func (p projectService) GetByID(ctx context.Context, id int64) (model.Project, error) {
	return model.NewProjectModel(p.db).First(query.Builder().Where("id", id))
}
