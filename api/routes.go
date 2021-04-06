package api

import (
	"github.com/mylxsw/container"
	"github.com/mylxsw/glacier/web"

	"github.com/mylxsw/wizard-enhance/api/controller"
)

func controllers(cc container.Resolver) []web.Controller {
	return []web.Controller{
		controller.NewWelcomeController(cc),
	}
}
