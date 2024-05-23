package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "Hello, world!\n")
}

func ArticleHandler(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "Article List\n")
}

func ArticleNumHandler(w http.ResponseWriter, req *http.Request)  {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func ArticleNiceHandler(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "Posting Nice\n")
}

func ArticleCommentHandler(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "Posting Comment\n")
}