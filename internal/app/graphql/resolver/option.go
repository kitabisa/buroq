package resolver

import (
	"github.com/kitabisa/buroq/internal/app/service"
)

type Options func(*Resolver)

func WithServices(svc *service.Services) Options {
	return func(r *Resolver) {
		r.svc = svc
	}
}
