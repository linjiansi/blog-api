package services

import "github.com/linjiansi/blog-api/models"

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}

type ArticleServicer interface {
	GetArticleService(articleID int) (models.Article, error)
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	PostFavoriteService(article models.Article) (models.Article, error)
}
