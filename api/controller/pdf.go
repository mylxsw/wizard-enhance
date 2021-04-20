package controller

import (
	"net/http"

	"github.com/mylxsw/glacier/infra"
	"github.com/mylxsw/glacier/web"
	"github.com/mylxsw/wizard-enhance/config"
	"github.com/mylxsw/wizard-enhance/pkg/pdf"
)

type PDFController struct {
	cc   infra.Resolver
	conf *config.Config
}

func NewPDFController(cc infra.Resolver, conf *config.Config) web.Controller {
	return &PDFController{cc: cc, conf: conf}
}

func (ctl PDFController) Register(router web.Router) {
	router.Group("pdf", func(router web.Router) {
		router.Post("office/", ctl.OfficeToPDF)
	})
}

func (ctl PDFController) OfficeToPDF(ctx web.Context, req web.Request) web.Response {
	source := req.Input("source")

	dstPath, err := pdf.OfficeToPDF(ctl.conf.GotenbergServer, ctl.conf.StoragePath, source)
	if err != nil {
		return ctx.JSONError(err.Error(), http.StatusInternalServerError)
	}

	return ctx.JSON(web.M{
		"path": dstPath,
	})
}

