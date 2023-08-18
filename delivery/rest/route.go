package rest

import (
	"article-api/delivery/rest/command"
	"article-api/delivery/rest/query"
	"article-api/libs/httpresponse"
	"net/http"

	"article-api/service"

	"github.com/labstack/echo/v4"
)

type Rest struct {
	goEnv   string
	query   query.QueryHandler
	command command.CommandHandler
}

// New ...
func New(
	svc service.Service,
	goEnv string,
) *Rest {
	return &Rest{
		goEnv:   goEnv,
		query:   query.NewQueryHandler(svc),
		command: command.NewCommandHandler(svc),
	}
}

func (re *Rest) Route(e *echo.Group) {
	ping := e.Group("/ping")
	ping.Any("", re.Ping)

	article := e.Group("/articles")
	article.GET("", re.query.ArticleHandler)
	article.POST("", re.command.CreateArticleHandler)
}

func (re *Rest) Ping(c echo.Context) error {
	// userID, ok := c.Get("user_id").(string)
	// err := errors.New("masokkkk")
	// if err != nil {
	// 	Error := errs.AssignErr(errs.AddTrace(err), errs.InvalidRequest)
	// 	httpresponse.ErrorTrace(c, http.StatusBadRequest, Error)
	// 	return nil
	// }
	httpresponse.Success(c, http.StatusOK, "events-api:"+re.goEnv)
	return nil
}
