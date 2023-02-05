package services

import "github.com/linjiansi/blog-api/models"

type BlogServicer interface {
	GetArticleService(articleID int) (models.Article, error)
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	PostFavoriteService(article models.Article) (models.Article, error)
	PostCommentService(comment models.Comment) (models.Comment, error)
}
