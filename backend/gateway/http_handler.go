package main

import (
	"net/http"

	"github.com/TA-Solaris/common"
	pb "github.com/TA-Solaris/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	o, err := h.meidaServiceClient.Upload(r.Context(), &pb.UploadRequest{
		UserID: userID,
	})
	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() == codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}

		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSON(w, http.StatusOK, o) // TODO Write media type (or maybe just a representation)?
}
