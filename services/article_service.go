package services

import (
	"github.com/linjiansi/blog-api/models"
	"github.com/linjiansi/blog-api/repositories"
)

func (s *BlogService) GetArticleService(articleID int) (models.Article, error) {

	returnedArticle, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	returnedCommentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	returnedArticle.CommentList = append(returnedArticle.CommentList, returnedCommentList...)

	return returnedArticle, nil
}

func (s *BlogService) PostArticleService(article models.Article) (models.Article, error) {

	returnedArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	return returnedArticle, nil
}

func (s *BlogService) GetArticleListService(page int) ([]models.Article, error) {

	returnedArticleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}

	return returnedArticleList, nil
}

func (s *BlogService) PostFavoriteService(article models.Article) (models.Article, error) {

	err := repositories.UpdateFavoriteNum(s.db, article.ID)
	if err != nil {
		return models.Article{}, err
	}
	article.FavoriteNum += 1

	return article, nil
}
