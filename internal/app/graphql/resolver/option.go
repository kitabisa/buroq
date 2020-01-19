package resolver

import (
	"github.com/kitabisa/buroq/internal/app/commons"
	"github.com/kitabisa/buroq/internal/app/service"
)

type Options func(*Resolver)

func WithServices(svc *service.Services) Options {
	return func(r *Resolver) {
		r.svc = svc
	}
}

func WithAllOptions(opt commons.Options) Options {
	return func(r *Resolver) {
		r.opt = opt
	}
}
