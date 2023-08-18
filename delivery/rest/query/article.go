package query

import (
	"article-api/libs/httpresponse"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (re *rest) ArticleHandler(c echo.Context) error {
	ctx := c.Request().Context()
	query := c.QueryParam("query")
	author := c.QueryParam("author")

	res, err := re.svc.ArticleQuery.Article(ctx, query, author)
	if err != nil {
		httpresponse.Error(c, http.StatusInternalServerError, httpresponse.ErrMessage{
			ErrMapping: err,
			Message:    err.Error(),
		})
		return err
	}

	httpresponse.Success(c, http.StatusOK, res)
	return nil
}
