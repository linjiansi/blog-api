package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/linjiansi/blog-api/controllers"
)

func NewRouter(con *controllers.BlogController) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/article", func(r chi.Router) {
		r.Post("/", con.PostArticleHandler)
		r.Post("/favorite", con.PostFavoriteArticleHandler)

		r.Get("/list", con.GetArticleListHandler)

		r.Route("/{articleID:[0-9]+}", func(r chi.Router) {
			r.Get("/", con.GetArticleDetailHandler)
		})
	})

	r.Post("/comment", con.PostCommentHandler)

	return r
}
