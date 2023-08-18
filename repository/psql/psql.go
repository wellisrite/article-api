package psql

import (
	"article-api/domain/article"
	"article-api/domain/orm"
	etPS "article-api/repository/psql/article"
	corm "article-api/repository/psql/orm"

	"github.com/jinzhu/gorm"
	gormV2 "gorm.io/gorm"
)

type PsqlRepositories struct {
	DbV2    *gormV2.DB
	Orm     orm.Repository
	Article article.Repository
}

func NewPsqlRepositories(db *gorm.DB, dbV2 *gormV2.DB) *PsqlRepositories {
	Orm := corm.New(db)
	return &PsqlRepositories{
		DbV2:    dbV2,
		Orm:     Orm,
		Article: etPS.New(Orm, dbV2),
	}
}
