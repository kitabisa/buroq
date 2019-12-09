package handler

import (
	"github.com/kitabisa/buroq/internal/app/commons"
	"github.com/kitabisa/buroq/internal/app/service"
)

type HandlerOption struct {
	commons.Options
	*service.Services
}
