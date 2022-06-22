package handlers

import (
	"encoding/json"
	"net/http"

	user "github.com/Questee29/taxi-app_driverService/models/driver"
	_ "github.com/lib/pq"
)

type UsersSignupService interface {
	IsRegistred(email, phone string) (bool, error)
	IsPasswordValid(password string) bool
	IsEmailValid(email string) bool
	IsNumberValid(number string) bool
	IsTaxiTypeValid(taxiType string) bool
	RegisterUser(name, phone, email, password, taxiType string) error
}

type SignupHandler struct {
	service UsersSignupService
}

func NewSignup(u UsersSignupService) *SignupHandler {
	return &SignupHandler{
		service: u,
	}
}

func (handler *SignupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u user.Driver

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !handler.service.IsPasswordValid(u.Password) {
		http.Error(w, "Bad password,try again. At least 8 chars(at least 1 upper,1 lower,1 num)", http.StatusBadRequest)
		return
	}

	if !handler.service.IsEmailValid(u.Email) {
		http.Error(w, "invalid email. Example : example@gmail.com", http.StatusBadRequest)
		return
	}

	if !handler.service.IsNumberValid(u.Phone) {
		http.Error(w, "invalid phone number. Only for belarus users", http.StatusBadRequest)
		return
	}

	if !handler.service.IsTaxiTypeValid(u.TaxiType) {
		http.Error(w, "invalid taxi type. There are 3 types : Economy, Comfort, Business", http.StatusBadRequest)
		return
	}

	isRegistred, err := handler.service.IsRegistred(u.Email, u.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isRegistred {
		http.Error(w, "Sorry, email or phone number already exists", http.StatusBadRequest)
		return
	}

	if err := handler.service.RegisterUser(u.Name, u.Phone, u.Email, u.Password, u.TaxiType); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
