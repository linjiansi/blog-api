package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/linjiansi/blog-api/controllers"
)

func NewRouter(
	aCon *controllers.ArticleController,
	cCon *controllers.CommentController,
) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/article", func(r chi.Router) {
		r.Post("/", aCon.PostArticleHandler)
		r.Post("/favorite", aCon.PostFavoriteArticleHandler)

		r.Get("/list", aCon.GetArticleListHandler)

		r.Route("/{articleID:[0-9]+}", func(r chi.Router) {
			r.Get("/", aCon.GetArticleDetailHandler)
		})
	})

	r.Post("/comment", cCon.PostCommentHandler)

	return r
}
