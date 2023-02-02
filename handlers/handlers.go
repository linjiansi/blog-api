package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world! \n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article... \n")
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)
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
