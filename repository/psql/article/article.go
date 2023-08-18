package article

import (
	"context"

	"article-api/domain/article"

	"gorm.io/gorm"
)

func (p *PostgreSQL) GetArticle(ctx context.Context, query, author string) (res []article.Article, err error) {
	// Articles not found in cache, fetch from the database using GORM
	queryBuilder := p.DBV2.Order("created_at DESC") // Sort by newest first

	if query != "" {
		// Apply search query to title and body
		queryBuilder = queryBuilder.Where("title LIKE ? OR body LIKE ?", "%"+query+"%", "%"+query+"%")
	}
	if author != "" {
		// Filter by author name
		queryBuilder = queryBuilder.Where("author = ?", author)
	}

	err = queryBuilder.Find(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, gorm.ErrRecordNotFound // You can define your own error
		}
		return res, err
	}

	return res, err
}
