package services

import (
	"github.com/linjiansi/blog-api/models"
	"github.com/linjiansi/blog-api/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	db.Close()

	returnedComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return returnedComment, err
}
