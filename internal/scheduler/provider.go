package scheduler

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/mylxsw/asteria/log"
	"github.com/mylxsw/glacier/infra"
	"github.com/mylxsw/glacier/scheduler"
	"github.com/mylxsw/go-utils/file"
	"github.com/mylxsw/wizard-enhance/config"
	"github.com/mylxsw/wizard-enhance/internal/service"
	"github.com/mylxsw/wizard-enhance/pkg/pdf"
	"gopkg.in/guregu/null.v3"
)

type Provider struct{}

func (s Provider) Aggregates() []infra.Provider {
	return []infra.Provider{
		scheduler.Provider(s.jobCreator),
	}
}

func (s Provider) Register(cc infra.Binder) {}
func (s Provider) Boot(cc infra.Resolver)   {}

func (s Provider) jobCreator(cc infra.Resolver, creator scheduler.JobCreator) {
	creator.Add("attachment-migrate", "@every 60s", scheduler.WithoutOverlap(func(attaSrv service.AttachmentService) {
		attachments, err := attaSrv.NeedGenerateExt(context.TODO())
		if err != nil {
			log.Errorf("query attachments need migrate failed: %v", err)
			return
		}

		for _, atta := range attachments {
			atta.FileType = null.StringFrom(strings.TrimPrefix(filepath.Ext(atta.Path.ValueOrZero()), "."))
			_ = atta.Save()
		}
	}))
	creator.Add("pdf-generator", "@every 30s", scheduler.WithoutOverlap(func(conf *config.Config, attaSrv service.AttachmentService) {
		attachments, err := attaSrv.NeedTransforms(context.TODO(), "pdf", "txt", "rtf", "fodt", "doc", "docx", "odt", "xls", "xlsx", "ods", "ppt", "pptx", "odp")
		if err != nil {
			log.Errorf("query attachments need transforms failed: %v", err)
			return
		}

		for _, atta := range attachments {
			if !file.Exist(filepath.Join(conf.StoragePath, atta.Path.ValueOrZero())) {
				log.Errorf("file not exist: %s", filepath.Join(conf.StoragePath, atta.Path.ValueOrZero()))
				atta.PreviewPath = null.StringFrom("")
				_ = atta.Save()
				continue
			}

			dst, err := pdf.OfficeToPDF(conf.GotenbergServer, conf.StoragePath, atta.Path.ValueOrZero())
			if err != nil {
				log.Errorf("pdf-generator failed: %v", err)
				continue
			}

			atta.PreviewPath = null.StringFrom(dst)
			if err := atta.Save(); err != nil {
				log.Errorf("save attachment failed: %v", err)
				continue
			}

			log.Debugf("pdf-generator: %s", dst)
		}
	}))
}
