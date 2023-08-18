package command

import (
	"article-api/service"

	"github.com/labstack/echo/v4"
)

type CommandHandler interface {
	CreateArticleHandler(c echo.Context) error
}

type (
	rest struct {
		svc service.Service
	}
)

func NewCommandHandler(
	svc service.Service,
) *rest {
	return &rest{
		svc: svc,
	}
}
