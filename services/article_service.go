package services

import (
	"github.com/linjiansi/blog-api/models"
	"github.com/linjiansi/blog-api/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	returnedArticle, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	returnedCommentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	returnedArticle.CommentList = append(returnedArticle.CommentList, returnedCommentList...)

	return returnedArticle, nil
}

func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	returnedArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return returnedArticle, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	returnedArticleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return returnedArticleList, nil
}

func PostFavoriteService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	err = repositories.UpdateFavoriteNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}
	article.FavoriteNum += 1

	return article, nil
}
