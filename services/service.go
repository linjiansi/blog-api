package services

import "database/sql"

type BlogService struct {
	db *sql.DB
}

func NewBlogService(db *sql.DB) *BlogService {
	return &BlogService{db: db}
}
