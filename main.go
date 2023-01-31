package main

import (
	"github.com/linjiansi/blog-api/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.GetArticleListHandler)
	http.HandleFunc("/article/1", handlers.GetArticleDetailHandler)
	http.HandleFunc("/article/favorite", handlers.PostFavoriteArticleHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
