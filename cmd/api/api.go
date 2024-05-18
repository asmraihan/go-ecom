package api

import (
	"database/sql"
	"go-ecom/service/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServerType struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServerType {
	return &APIServerType{
		addr: addr,
		db:   db,
	}
}

func (server *APIServerType) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userService := user.NewHandler()
	userService.UserRoutes(subRouter)

	log.Println("Listening server on", server.addr)
	return http.ListenAndServe(server.addr, router)
}
