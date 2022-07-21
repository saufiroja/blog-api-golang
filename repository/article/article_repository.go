package article

import "echo/blog-api/entity"

type ArticleRepository interface {
	CreateArticle(article entity.Article) (entity.Article, error)
}
