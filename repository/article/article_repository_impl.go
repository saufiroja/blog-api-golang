package article

import (
	"echo/blog-api/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateArticle(article entity.Article) (entity.Article, error) {
	err := r.db.Create(&article).Error
	return article, err
}

func (r *repository) FindAllArticle() ([]entity.Article, error) {
	var articles []entity.Article
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *repository) FindByIDArticle(id string) (entity.Article, error) {
	var article entity.Article
	err := r.db.Where("id = ?", id).First(&article).Error
	return article, err
}

func (r *repository) UpdateArticle(id string, article entity.Article) (entity.Article, error) {
	err := r.db.Model(&article).Where("id = ?", id).Updates(&article).Error
	return article, err
}

func (r *repository) DeleteArticle(id string, article entity.Article) error {
	err := r.db.Where("id = ?", id).Delete(&entity.Article{}).Error
	return err
}
