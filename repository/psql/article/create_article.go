package article

import (
	"context"

	"article-api/domain/article"
)

func (p *PostgreSQL) CreateArticle(ctx context.Context, article article.Article) (err error) {
	// Insert the new article into the database using GORM
	err = p.DBV2.Create(&article).Error
	if err != nil {
		// Handle the error appropriately, possibly return an error or log it
		return err
	}

	return
}
