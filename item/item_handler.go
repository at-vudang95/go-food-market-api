package item

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

// GetAll api
func (h *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
}

// GetCart api
func (h *HTTPHandler) GetCart(w http.ResponseWriter, r *http.Request) {
}

// GetItemByName api
func (h *HTTPHandler) GetItemByName(w http.ResponseWriter, r *http.Request) {
}

// GetItemByCategory api
func (h *HTTPHandler) GetItemByCategory(w http.ResponseWriter, r *http.Request) {
}

// GetItemByID api
func (h *HTTPHandler) GetItemByID(w http.ResponseWriter, r *http.Request) {
}

// GetItemNews api
func (h *HTTPHandler) GetItemNews(w http.ResponseWriter, r *http.Request) {
}

// GetItemPromotion api
func (h *HTTPHandler) GetItemPromotion(w http.ResponseWriter, r *http.Request) {
}

// GetItemBestSale api
func (h *HTTPHandler) GetItemBestSale(w http.ResponseWriter, r *http.Request) {
}

// GetItemSearch api
func (h *HTTPHandler) GetItemSearch(w http.ResponseWriter, r *http.Request) {
}

// GetCategories api
func (h *HTTPHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
}

// GetItemBySupplier api
func (h *HTTPHandler) GetItemBySupplier(w http.ResponseWriter, r *http.Request) {
}

// GetItemSearchByCategory api
func (h *HTTPHandler) GetItemSearchByCategory(w http.ResponseWriter, r *http.Request) {
}

// GetItemSearchByStatus api
func (h *HTTPHandler) GetItemSearchByStatus(w http.ResponseWriter, r *http.Request) {
}

// PostItem api
func (h *HTTPHandler) PostItem(w http.ResponseWriter, r *http.Request) {
}

// PutItem api
func (h *HTTPHandler) PutItem(w http.ResponseWriter, r *http.Request) {
}

// DeleteItem api
func (h *HTTPHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository, s *infrastructure.SQL) *HTTPHandler {
	// outfit set.
	or := NewRepository(br, s.Master)
	ou := NewUsecase(bu, s.Master, or)
	return &HTTPHandler{BaseHTTPHandler: *bh, usecase: ou}
}
