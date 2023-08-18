package command

import (
	"article-api/domain/article"
	"article-api/libs/httpresponse"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (re *rest) CreateArticleHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var articleBody article.ArticleWrite
	if err := json.NewDecoder(c.Request().Body).Decode(&articleBody); err != nil {
		return c.String(http.StatusBadRequest, "Invalid JSON")
	}

	err := re.svc.ArticleCommand.CreateArticle(ctx, articleBody)
	if err != nil {
		httpresponse.Error(c, http.StatusInternalServerError, httpresponse.ErrMessage{
			ErrMapping: err,
			Message:    err.Error(),
		})
		return err
	}

	httpresponse.Success(c, http.StatusCreated, nil)
	return nil
}
