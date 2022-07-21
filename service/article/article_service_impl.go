package article

import (
	"echo/blog-api/config"
	"echo/blog-api/entity"
	"echo/blog-api/repository/article"
)

type Service struct {
	r    article.ArticleRepository
	conf config.Config
}

func NewArticleService(r article.ArticleRepository, conf config.Config) ArticleService {
	return &Service{r: r, conf: conf}
}

func (s *Service) CreateArticle(article entity.Article) (entity.Article, error) {
	return s.r.CreateArticle(article)
}
