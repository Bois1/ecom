package user

import (
	"errors"
	"net/http"
	"time"
	"fmt"

	"github.com/Bois1/ecomm/service/auth"
	"github.com/Bois1/ecomm/types"
	"github.com/Bois1/ecomm/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := payload.Validate(); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	
	existingUser, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("database error: %w", err))
		return
	}
	if existingUser != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("user with this email already exists"))
		return
	}

	
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to hash password: %w", err))
		return
	}

	
	user := &types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	if err := h.store.CreateUser(user); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create user: %w", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, user)
}