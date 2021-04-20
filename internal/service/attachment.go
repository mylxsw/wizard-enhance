package service

import (
	"context"
	"database/sql"

	"github.com/mylxsw/eloquent/query"
	"github.com/mylxsw/glacier/infra"
	"github.com/mylxsw/wizard-enhance/internal/service/model"
)

type AttachmentService interface {
	GetByID(ctx context.Context, id int64) (model.Attachment, error)
	All(ctx context.Context) ([]model.Attachment, error)
	NeedTransforms(ctx context.Context, ext ...interface{}) ([]model.Attachment, error)
	NeedGenerateExt(ctx context.Context) ([]model.Attachment, error)
}

func NewAttachmentService(cc infra.Resolver, db *sql.DB) AttachmentService {
	return &attachmentService{db: db, cc: cc}
}

type attachmentService struct {
	db *sql.DB
	cc infra.Resolver
}

func (p attachmentService) All(ctx context.Context) ([]model.Attachment, error) {
	return model.NewAttachmentModel(p.db).Get()
}

func (p attachmentService) NeedTransforms(ctx context.Context, ext ...interface{}) ([]model.Attachment, error) {
	return model.NewAttachmentModel(p.db).Get(query.Builder().WhereIn("file_type", ext...).WhereNull("preview_path"))
}

func (p attachmentService) GetByID(ctx context.Context, id int64) (model.Attachment, error) {
	return model.NewAttachmentModel(p.db).First(query.Builder().Where("id", id))
}

func (p attachmentService) NeedGenerateExt(ctx context.Context) ([]model.Attachment, error) {
	return model.NewAttachmentModel(p.db).Get(query.Builder().WhereNull("file_type"))
}
