package service

import (
	"article-api/repository/psql"
	command "article-api/service/command"
	query "article-api/service/query"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

type Service struct {
	ArticleQuery   query.ArticleQueryService
	ArticleCommand command.ArticleCommandService
}

func New(
	validator echo.Validator,
	psqlRepos *psql.PsqlRepositories,
	redis *redis.Client,
) Service {
	svc := Service{
		ArticleQuery:   *query.NewArticleQueryService(validator, redis, psqlRepos),
		ArticleCommand: *command.NewArticleCommandService(validator, redis, psqlRepos),
	}

	return svc
}
