package query

import (
	"article-api/domain/article"
	"article-api/repository/psql"
	"context"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

type ArticleQueryServiceInterface interface {
	Article(ctx context.Context, id string) (res article.Article, err error)
}

type ArticleQueryService struct {
	validator   echo.Validator
	redisClient *redis.Client
	psqlRepo    *psql.PsqlRepositories
}

func NewArticleQueryService(validator echo.Validator, redisClient *redis.Client, psqlRepo *psql.PsqlRepositories) *ArticleQueryService {
	return &ArticleQueryService{
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
