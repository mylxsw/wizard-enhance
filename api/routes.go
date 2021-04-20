package api

import (
	"github.com/mylxsw/container"
	"github.com/mylxsw/glacier/web"
	"github.com/mylxsw/wizard-enhance/config"

	"github.com/mylxsw/wizard-enhance/api/controller"
)

func controllers(cc container.Resolver, conf *config.Config) []web.Controller {
	return []web.Controller{
		controller.NewInspectController(cc, conf),
		controller.NewPDFController(cc, conf),
	}
}
