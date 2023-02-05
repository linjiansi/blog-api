package api

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/linjiansi/blog-api/api/middlewares"
	"github.com/linjiansi/blog-api/controllers"
	"github.com/linjiansi/blog-api/services"
)

func NewRouter(db *sql.DB) *chi.Mux {

	s := services.NewBlogService(db)
	aCon := controllers.NewArticleController(s)
	cCon := controllers.NewCommentController(s)

	r := chi.NewRouter()
	r.Use(
		middlewares.LoggingMiddleware,
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
	)

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
