package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/linjiansi/blog-api/handlers"
	"log"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/hello", handlers.HelloHandler)
	r.Post("/article", handlers.PostArticleHandler)
	r.Get("/article/list", handlers.GetArticleListHandler)
	r.Route("/article", func(r chi.Router) {
		r.Route("/{articleID:[0-9]+}", func(r chi.Router) {
			r.Get("/", handlers.GetArticleDetailHandler)
		})
	})

	r.Post("/article/favorite", handlers.PostFavoriteArticleHandler)
	r.Post("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
