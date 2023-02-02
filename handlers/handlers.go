package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world! \n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article... \n")
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List \n")
}

func GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := chi.URLParam(req, "articleID")
	resString := fmt.Sprintf("Article No.%s\n", articleID)
	io.WriteString(w, resString)
}

func PostFavoriteArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice... \n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment... \n")
}
