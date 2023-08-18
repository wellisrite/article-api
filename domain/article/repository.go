package article

import (
	"context"
)

type Repository interface {
	GetArticle(ctx context.Context, query, author string) (res []Article, err error)
	CreateArticle(ctx context.Context, article Article) (err error)
}
