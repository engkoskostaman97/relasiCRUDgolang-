package routes

import (
  "dumbmerch/handlers"
  "dumbmerch/pkg/mysql"
  "dumbmerch/repositories"

  "github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
  userRepository := repositories.RepositoryUser(mysql.DB)
  h := handlers.HandlerAuth(userRepository)

  r.HandleFunc("/register", h.Register).Methods("POST")
  r.HandleFunc("/login", h.Login).Methods("POST") // add this code
}