package handlers

import "net/http"

type Handler struct {
}

func NewHandler() (*Handler, error) {

	return &Handler{}, nil
}
func LoadUnfinishedWork(w http.ResponseWriter, r *http.Request) {

}

func SaveNewUrls(w http.ResponseWriter, r *http.Request) {

}

func LoadUrls(w http.ResponseWriter, r *http.Request) {

}
