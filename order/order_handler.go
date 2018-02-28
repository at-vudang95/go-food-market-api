package order

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

// GetOrderByUser api
func (h *HTTPHandler) GetOrderByUser(w http.ResponseWriter, r *http.Request) {
}

// GetOrderByID api
func (h *HTTPHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
}

// GetOrderByStatus api
func (h *HTTPHandler) GetOrderByStatus(w http.ResponseWriter, r *http.Request) {
}

// DeleteOrderItem api
func (h *HTTPHandler) DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
}

// PostOrder api
func (h *HTTPHandler) PostOrder(w http.ResponseWriter, r *http.Request) {
}

// PutOrder api
func (h *HTTPHandler) PutOrder(w http.ResponseWriter, r *http.Request) {
}

// DeleteOrder api
func (h *HTTPHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository, s *infrastructure.SQL) *HTTPHandler {
	// outfit set.
	or := NewRepository(br, s.Master)
	ou := NewUsecase(bu, s.Master, or)
	return &HTTPHandler{BaseHTTPHandler: *bh, usecase: ou}
}
