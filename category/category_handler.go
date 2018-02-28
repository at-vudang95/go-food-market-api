package category

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

// GetCategoryByID api
func (h *HTTPHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
}

// GetCategoryByParentID api
func (h *HTTPHandler) GetCategoryByParentID(w http.ResponseWriter, r *http.Request) {
}

// GetCategoryByLevel api
func (h *HTTPHandler) GetCategoryByLevel(w http.ResponseWriter, r *http.Request) {
}

// PostCategory api
func (h *HTTPHandler) PostCategory(w http.ResponseWriter, r *http.Request) {
}

// PutCategory api
func (h *HTTPHandler) PutCategory(w http.ResponseWriter, r *http.Request) {
}

// DeleteCategory api
func (h *HTTPHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository, s *infrastructure.SQL) *HTTPHandler {
	// outfit set.
	or := NewRepository(br, s.Master)
	ou := NewUsecase(bu, s.Master, or)
	return &HTTPHandler{BaseHTTPHandler: *bh, usecase: ou}
}
