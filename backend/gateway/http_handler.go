package main

import (
	"net/http"

	pb "github.com/TA-Solaris/common/api"
)

type handler struct {
	meidaServiceClient pb.MediaServiceClient
}

func NewHandler(meidaServiceClient pb.MediaServiceClient) *handler {
	return &handler{meidaServiceClient}
}

func (h *handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/user/{userID}/media", h.HandleMediaUpload)
}

func (h *handler) HandleMediaUpload(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")

	// TODO Get media

	h.meidaServiceClient.Upload(r.Context(), &pb.UploadRequest{
		UserID: userID,
	})
}
