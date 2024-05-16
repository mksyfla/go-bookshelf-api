package main

import (
	"bookshelf-api-pemula-dicoding/handler"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.POST("/books", handler.PostBookHandler)
	router.GET("/books", handler.GetBooksHandler)
	router.GET("/books/:id", handler.GetBookByIdHandler)
	router.PUT("/books/:id", handler.PutBookByIdHandler)
	router.DELETE("/books/:id", handler.DeleteBookByIdHandler)

	server := http.Server{
		Addr:    "localhost:9000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
