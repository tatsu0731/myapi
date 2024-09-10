package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tatsu0731/myapi/handlers"
)

func main()  {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleNumHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.ArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.ArticleCommentHandler).Methods(http.MethodPost)

	log.Println("Server start as port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}