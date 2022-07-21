package article

import "echo/blog-api/entity"

type ArticleRepository interface {
	CreateArticle(article entity.Article) (entity.Article, error)
	FindAllArticle() ([]entity.Article, error)
	FindByIDArticle(id string) (entity.Article, error)
	UpdateArticle(id string, article entity.Article) (entity.Article, error)
	DeleteArticle(id string, article entity.Article) error
}
