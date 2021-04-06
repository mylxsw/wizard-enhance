package service

import (
	"github.com/mylxsw/glacier/infra"
)

type Provider struct{}

func (p Provider) Register(cc infra.Binder) {
	cc.MustSingletonOverride(NewProjectService)
}

func (p Provider) Boot(cc infra.Resolver) {
}
