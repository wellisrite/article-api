package query

import (
	"article-api/service"

	"github.com/labstack/echo/v4"
)

type QueryHandler interface {
	ArticleHandler(c echo.Context) error
}

type (
	rest struct {
		svc service.Service
	}
)

func NewQueryHandler(
	svc service.Service,
) *rest {
	return &rest{
		svc: svc,
	}
}
