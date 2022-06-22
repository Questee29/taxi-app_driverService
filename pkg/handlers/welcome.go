package handlers

import (
	"fmt"
	"net/http"

	"github.com/Questee29/taxi-app_driverService/middleware"
)

type UsersAuthService interface {
	ParseToken(tokenString string) (string, error)
	DeleteToken(email string) error
}

type Handler struct {
	usersService UsersAuthService
}

func NewWelcome(usersService UsersAuthService) *Handler {
	return &Handler{
		usersService: usersService,
	}
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value(middleware.ContextUserKey)
	w.Write([]byte(fmt.Sprintf("Welcome driver %s!", userEmail)))

}
