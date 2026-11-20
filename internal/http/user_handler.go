// Package httpapi
package httpapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	userapp "uzpay/internal/application/user"
	"uzpay/internal/domain/user"
)

type registerRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

type registerResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type UserHandler struct {
	svc *userapp.Service
}

func NewUserHandler(svc *userapp.Service) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	u, err := h.svc.Register(r.Context(), req.Username, req.Email, req.Password, req.PhoneNumber)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrEmailTaken):
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			log.Println("Internal error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	res := registerResponse{
		ID:          u.ID().String(),
		Username:    u.Username(),
		Email:       u.Email(),
		PhoneNumber: u.PhoneNumber(),
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("encode register response", err)
	}
}
