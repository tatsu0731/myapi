package main

import (
	"log"
	"net/http"
	"example.com/api/models"
)

func main()  {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.ArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleNumHandler)
	http.HandleFunc("/article/nice", handlers.ArticleNiceHandler)
	http.HandleFunc("/comment", handlers.ArticleCommentHandler)

	log.Println("Server start as port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}