package handlers

import (
	"encoding/json"
	"net/http"

	user "github.com/Questee29/taxi-app_driverService/models/driver"
)

type UsersSignInService interface {
	GenerateJWT(number, password string) (string, error)
}

type SignInHandler struct {
	usersService UsersSignInService
}

func NewSignIn(usersSignInService UsersSignInService) *SignInHandler {
	return &SignInHandler{
		usersService: usersSignInService,
	}
}

func (handler *SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input user.AuthDetails
	//reads json from user
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//check from database
	token, err := handler.usersService.GenerateJWT(input.Phone, input.Password)
	if err != nil {
		http.Error(w, "wrong number or pass,try again", http.StatusInternalServerError)
		return
	}

	//generates newJWT and set Cookie

	//just for return cookies json
	if err := json.NewEncoder(w).Encode(&token); err != nil {
		http.Error(w, "sorry", http.StatusInternalServerError)
	}
}
