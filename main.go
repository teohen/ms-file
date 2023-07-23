package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	fmt.Println("vim-go")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello wordl"))
	})
	router.Post("/file", uploadFile)

	http.ListenAndServe(":3000", router)

}

func uploadFile(w http.ResponseWriter, r *http.Request) {

	type Result struct {
		Id string `json:"id"`
	}

	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("file")

	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
	}

	defer file.Close()

	tempFile, err := ioutil.TempFile("data", "upload-*.png")

	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	h := sha256.New()

	h.Write(fileBytes)

	bs := h.Sum(nil)

	hash := fmt.Sprintf("%x", string(bs))

	tempFile.Write(fileBytes)

	res := Result{}
	res.Id = hash

	data, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
