package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Sahil2k07/go-SOLID/src/middlewares"
	"github.com/Sahil2k07/go-SOLID/src/services"
)

type UserRoutes interface {
	Signup(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	service *services.UserService
}

func UserControllerInit(s *services.UserService) *UserController {
	return &UserController{service: s}
}

func (uc *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	// logic

	json.NewEncoder(w).Encode("Signup successfull")
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	// logic

	json.NewEncoder(w).Encode("Login successfull")
}

func (uc *UserController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middlewares.UserContext).(*middlewares.UserData) // type assertion
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"data":    user,
	}

	json.NewEncoder(w).Encode(response)
}
