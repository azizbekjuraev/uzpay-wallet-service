// Package httpapi
package httpapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	walletapp "uzpay/internal/application/wallet"
	"uzpay/internal/domain/wallet"
)

type registerWalletRequest struct {
	UserID string `json:"userId"`
}

type registerWalletResponse struct {
	ID      string `json:"id"`
	Balance int64  `json:"balance"`
}

type WalletHandler struct {
	svc *walletapp.Service
}

func NewWalletHandler(svc *walletapp.Service) *WalletHandler {
	return &WalletHandler{svc: svc}
}

func (h *WalletHandler) Register(writer http.ResponseWriter, r *http.Request) {
	var req registerWalletRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(writer, "invalid json", http.StatusBadRequest)
		return
	}
	w, err := h.svc.Register(r.Context(), req.UserID)
	if err != nil {
		switch {
		case errors.Is(err, wallet.ErrWalletExists):
			http.Error(writer, err.Error(), http.StatusConflict)
		default:
			log.Println("Internal error")
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	res := registerWalletResponse{
		ID:      w.ID().String(),
		Balance: w.Balance(),
	}
	if err := json.NewEncoder(writer).Encode(res); err != nil {
		log.Println("encode register response", err)
	}
}
