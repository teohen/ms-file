package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	fmt.Println("vim-go")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/files", HandleUploadFile)
	router.Get("/files/{id}", HandleGetFile)
	http.ListenAndServe(":3000", router)
}
