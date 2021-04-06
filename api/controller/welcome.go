package controller

import (
	"github.com/mylxsw/glacier/infra"
	"github.com/mylxsw/glacier/web"

	"github.com/mylxsw/wizard-enhance/config"
)

type WelcomeController struct {
	cc   infra.Resolver
	conf *config.Config
}

func NewWelcomeController(cc infra.Resolver) web.Controller {
	conf := config.Get(cc)
	return &WelcomeController{cc: cc, conf: conf}
}

func (wel WelcomeController) Register(router web.Router) {
	router.Group("/home", func(router web.Router) {
		router.Any("/", wel.Home)
	})
}

func (wel WelcomeController) Home(ctx web.Context) web.Response {
	return ctx.JSON(web.M{
		"hello": "world",
	})
}
