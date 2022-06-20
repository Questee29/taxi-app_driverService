package handlers

import (
	"net/http"
)

type HandlerLogOut struct {
	usersService UsersAuthService
}

func NewLogOut(usersService UsersAuthService) *HandlerLogOut {
	return &HandlerLogOut{
		usersService: usersService,
	}
}

func (handler *HandlerLogOut) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
