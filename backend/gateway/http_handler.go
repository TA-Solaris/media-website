package main

import "net/http"

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/user/{userID}/media", h.HandleMediaUpload)
}

func (h *handler) HandleMediaUpload(w http.ResponseWriter, r *http.Request) {

}
