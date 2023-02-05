package services

import (
	"database/sql"
	"errors"
	"github.com/linjiansi/blog-api/apperrors"
	"github.com/linjiansi/blog-api/models"
	"github.com/linjiansi/blog-api/repositories"
)

func (s *BlogService) GetArticleService(articleID int) (models.Article, error) {

	returnedArticle, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		} else {
			err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
			return models.Article{}, err
		}
	}

	returnedCommentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	returnedArticle.CommentList = append(returnedArticle.CommentList, returnedCommentList...)

	return returnedArticle, nil
}

func (s *BlogService) PostArticleService(article models.Article) (models.Article, error) {

	returnedArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to recode data")
		return models.Article{}, err
	}

	return returnedArticle, nil
}

func (s *BlogService) GetArticleListService(page int) ([]models.Article, error) {

	returnedArticleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(returnedArticleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return returnedArticleList, nil
}

func (s *BlogService) PostFavoriteService(article models.Article) (models.Article, error) {

	err := repositories.UpdateFavoriteNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "not exist target article")
			return models.Article{}, err
		} else {
			err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
			return models.Article{}, err
		}

	}
	article.FavoriteNum += 1

	return article, nil
}
