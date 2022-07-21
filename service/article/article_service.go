package article

import "echo/blog-api/entity"

type ArticleService interface {
	CreateArticle(article entity.Article) (entity.Article, error)
}
