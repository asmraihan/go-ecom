package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerType struct {
}

func NewHandler() *HandlerType {
	return &HandlerType{}
}

// handlers for user routes
func (handler *HandlerType) UserRoutes(router *mux.Router) {
	router.HandleFunc("/login", handler.handleLogin).Methods("POST")
	router.HandleFunc("/register", handler.handleRegister).Methods("POST")
}

func (handler *HandlerType) handleLogin(w http.ResponseWriter, r *http.Request) {

}
func (handler *HandlerType) handleRegister(w http.ResponseWriter, r *http.Request) {

}
