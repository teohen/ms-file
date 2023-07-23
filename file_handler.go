package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type Result struct {
	Id string `json:"id"`
}

func getHash(fileBytes []byte) string {
	h := sha256.New()

	h.Write(fileBytes)

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", string(bs))
}

func printError(error error, msg string) {
	if error != nil {
		fmt.Println(msg)
		fmt.Println(error)
	}
}

func HandleUploadFile(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("file")

	printError(err, "error retrieving file from form")

	defer file.Close()

	hash := saveFile(file)

	res := Result{
		Id: hash,
	}

	data, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func getFileBytesFrom(file multipart.File) []byte {
	var fileBytes []byte
	fileBytes, err := ioutil.ReadAll(file)

	printError(err, "error creating the filebytes")

	return fileBytes
}

func saveFile(file multipart.File) string {
	fileBytes := getFileBytesFrom(file)

	hash := getHash(fileBytes)

	err := os.WriteFile(fmt.Sprintf("data/%s.png", hash), fileBytes, 0666)

	printError(err, "error saving file on file system")
	return hash
}

func getFile(hash string) []byte {
	var file []byte
	file, err := os.ReadFile(fmt.Sprintf("data/%s.png", hash))

	if err != nil {
		printError(err, "could read file from file system")
	}
	fmt.Println("LENG", len(file))
	return file
}

func HandleGetFile(w http.ResponseWriter, r *http.Request) {
	status := 200
	fileHash := chi.URLParam(r, "id")
	file := getFile(fileHash)

	if len(file) < 1 {
		status = 404
	}

	w.Header().Add("application", "octet-stream")
	w.WriteHeader(status)
	w.Write(file)
}
