package handler

import (
	"github.com/kitabisa/buroq/internal/app/commons"
	"github.com/kitabisa/buroq/internal/app/service"
)

// HandlerOption option for handler, including all service
type HandlerOption struct {
	commons.Options
	*service.Services
}
