package article

import (
	corm "article-api/domain/orm"

	"gorm.io/gorm"
)

type PostgreSQL struct {
	Ormr corm.Repository
	DBV2 *gorm.DB
}

func New(ormr corm.Repository, dbv2 *gorm.DB) *PostgreSQL {
	return &PostgreSQL{ormr, dbv2}
}

const (
	TableArticle = "article"
)
