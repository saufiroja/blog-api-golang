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

func (s *Service) FindAllArticle() ([]entity.Article, error) {
	article, err := s.r.FindAllArticle()
	return article, err
}

func (s *Service) FindByIDArticle(id string) (entity.Article, error) {
	article, err := s.r.FindByIDArticle(id)
	return article, err
}

func (s *Service) UpdateArticle(id string, article entity.Article) (entity.Article, error) {
	// check if article exist
	_, err := s.r.FindByIDArticle(id)
	if err != nil {
		return article, err
	}
	// update article
	return s.r.UpdateArticle(id, article)
}

func (s *Service) DeleteArticle(id string, article entity.Article) error {
	// check if article exist
	_, err := s.r.FindByIDArticle(id)
	if err != nil {
		return err
	}
	// delete article
	return s.r.DeleteArticle(id, article)
}
