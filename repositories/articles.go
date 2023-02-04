package repositories

import (
	"database/sql"
	"github.com/linjiansi/blog-api/models"

	_ "github.com/go-sql-driver/mysql"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles(title, contents, username, favorite, created_at)
		values (?, ?, ?, 0, now());
	`

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)

	if err != nil {
		return article, err
	}

	id, _ := result.LastInsertId()
	article.ID = int(id)

	return article, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, favorite
		from articles limit ? offset ?;
	`
	rows, err := db.Query(sqlStr, articleNumPerPage, (page-1)*articleNumPerPage)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	articleArray := make([]models.Article, 0)

	for rows.Next() {
		var article models.Article
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Contents,
			&article.UserName,
			&article.FavoriteNum,
			&article.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = "select * from articles where article_id = ?;"

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	err := row.Scan(
		&article.ID,
		&article.Title,
		&article.Contents,
		&article.UserName,
		&article.FavoriteNum,
		&article.CreatedAt,
	)

	if err != nil {
		return models.Article{}, err
	}

	return article, err
}

func UpdateFavoriteNum(db *sql.DB, articleID int) error {
	const sqlGetFavorite = "select favorite from articles where article_id = ?;"
	const sqlUpdateFavorite = "update articles set favorite = ? where article_id = ?;"

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(sqlGetFavorite, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var favoriteNum int
	err = row.Scan(&favoriteNum)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateFavorite, favoriteNum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
