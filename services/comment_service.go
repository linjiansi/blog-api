package services

import (
	"github.com/linjiansi/blog-api/apperrors"
	"github.com/linjiansi/blog-api/models"
	"github.com/linjiansi/blog-api/repositories"
)

func (s *BlogService) PostCommentService(comment models.Comment) (models.Comment, error) {

	returnedComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return returnedComment, err
}
