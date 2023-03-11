package services

import (
	"database/sql"
	"errors"

	"github.com/linjiansi/blog-api/apperrors"
	"github.com/linjiansi/blog-api/models"
	"github.com/linjiansi/blog-api/repositories"
)

func (s *BlogService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	type articleResult struct {
		article models.Article
		err     error
	}

	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}

	articleCh := make(chan articleResult)
	commentCh := make(chan commentResult)
	defer close(articleCh)
	defer close(commentCh)

	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		article, articleGetErr = repositories.SelectArticleDetail(s.db, articleID)
		ch <- articleResult{article: article, err: articleGetErr}
	}(articleCh, s.db, articleID)

	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		commentList, commentGetErr = repositories.SelectCommentList(db, articleID)
		ch <- commentResult{commentList: &commentList, err: commentGetErr}
	}(commentCh, s.db, articleID)

	for i := 0; i < 2; i++ {
		select {
		case ar := <-articleCh:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentCh:
			commentList, commentGetErr = *cr.commentList, cr.err
		}
	}

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		} else {
			err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
			return models.Article{}, err
		}
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
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
