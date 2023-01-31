package handlers

import (
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
	io.WriteString(w, "Article No.1 \n")
}

func PostFavoriteArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice... \n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment... \n")
}
