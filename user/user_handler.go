package user

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

// Hello to register user ID which originates from Device ID.
//
// "First": Search User from Entity by Device ID.
// "Second": If User record exists,move to step "Finally".
// "Third": If User record does not exist, register device ID to Entity.
// "Finally":store User_ID acquired from Entity to JSON Web Token (JWT).
func (h *HTTPHandler) Hello(w http.ResponseWriter, r *http.Request) {
	common := CommonResponse{Message: "Parse request error.", Errors: nil}
	h.ResponseJSON(w, common)
}

// GetUserByID api
func (h *HTTPHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
}

// ResetPassword api
func (h *HTTPHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
}

// CheckEmail api
func (h *HTTPHandler) CheckEmail(w http.ResponseWriter, r *http.Request) {
}

// PostUser api
func (h *HTTPHandler) PostUser(w http.ResponseWriter, r *http.Request) {
}

// PutUser api
func (h *HTTPHandler) PutUser(w http.ResponseWriter, r *http.Request) {
}

// DeleteUser api
func (h *HTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository, s *infrastructure.SQL) *HTTPHandler {
	// user set.
	userRepo := NewRepository(br, s.Master, s.Read)
	userUsecase := NewUsecase(bu, s.Master, userRepo)
	return &HTTPHandler{BaseHTTPHandler: *bh, usecase: userUsecase}
}
