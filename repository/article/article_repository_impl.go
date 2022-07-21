package article

import (
	"echo/blog-api/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &repository{db: db}
}

func (r *repository) CreateArticle(article entity.Article) (entity.Article, error) {
	err := r.db.Create(&article).Error
	return article, err
}
