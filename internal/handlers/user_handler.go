package handlers

import (
	"auth-service/internal/services"
	"encoding/json"
	"net/http"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrMessage struct {
	Error string `json:"error"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errMsg := ErrMessage{
			Error: "Invalid input",
		}
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, errMsg.Error, http.StatusBadRequest)
		return
	}

	user, err := services.RegisterUser(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
