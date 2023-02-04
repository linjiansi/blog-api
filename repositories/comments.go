package repositories

import (
	"database/sql"
	"github.com/linjiansi/blog-api/models"

	_ "github.com/go-sql-driver/mysql"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = "insert into comments (article_id, message, created_at) values (?, ?, now());"

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message, comment.CreatedAt)

	if err != nil {
		return models.Comment{}, err
	}

	id, _ := result.LastInsertId()
	comment.CommentID = int(id)

	return comment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = "select * from comments where article_id = ?;"

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	commentArray := make([]models.Comment, 0)

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.CommentID,
			&comment.ArticleID,
			&comment.Message,
			&comment.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
