package command

import (
	"article-api/domain/article"
	"article-api/repository/psql"
	"context"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

type ArticleCommandServiceInterface interface {
	CreateArticle(ctx context.Context, id string) (res article.Article, err error)
}

type ArticleCommandService struct {
	validator   echo.Validator
	redisClient *redis.Client
	psqlRepo    *psql.PsqlRepositories
}

func NewArticleCommandService(validator echo.Validator, redisClient *redis.Client, psqlRepo *psql.PsqlRepositories) *ArticleCommandService {
	return &ArticleCommandService{
		validator:   validator,
		redisClient: redisClient,
		psqlRepo:    psqlRepo,
	}
}

const (
	TYPE_SEEN_NO_AT_ALL = "NO_AT_ALL"
	TYPE_SEEN_PCTG_ONLY = "PCTG_ONLY"
	TYPE_SEEN_ALL       = "ALL"
)
