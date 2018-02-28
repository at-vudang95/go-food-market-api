package image

import (
	"net/http"

	"github.com/at-vudang95/go-food-market-api/infrastructure"
	"github.com/at-vudang95/go-food-market-api/shared/handler"
	"github.com/at-vudang95/go-food-market-api/shared/repository"
	"github.com/at-vudang95/go-food-market-api/shared/usecase"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.BaseHTTPHandler
	usecase UsecaseInterface
}

// PostImage api
func (h *HTTPHandler) PostImage(w http.ResponseWriter, r *http.Request) {
}

// PutImage api
func (h *HTTPHandler) PutImage(w http.ResponseWriter, r *http.Request) {
}

// DeleteImage api
func (h *HTTPHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository, s *infrastructure.SQL) *HTTPHandler {
	// outfit set.
	or := NewRepository(br, s.Master)
	ou := NewUsecase(bu, s.Master, or)
	return &HTTPHandler{BaseHTTPHandler: *bh, usecase: ou}
}
