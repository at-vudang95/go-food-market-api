package promotion

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

// GetPromotionByID api
func (h *HTTPHandler) GetPromotionByID(w http.ResponseWriter, r *http.Request) {
}

// GetPromotions api
func (h *HTTPHandler) GetPromotions(w http.ResponseWriter, r *http.Request) {
}

// DeletePromotionItem api
func (h *HTTPHandler) DeletePromotionItem(w http.ResponseWriter, r *http.Request) {
}

// PostPromotion api
func (h *HTTPHandler) PostPromotion(w http.ResponseWriter, r *http.Request) {
}

// PutPromotion api
func (h *HTTPHandler) PutPromotion(w http.ResponseWriter, r *http.Request) {
}

// DeletePromotion api
func (h *HTTPHandler) DeletePromotion(w http.ResponseWriter, r *http.Request) {
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository, s *infrastructure.SQL) *HTTPHandler {
	// outfit set.
	or := NewRepository(br, s.Master)
	ou := NewUsecase(bu, s.Master, or)
	return &HTTPHandler{BaseHTTPHandler: *bh, usecase: ou}
}
