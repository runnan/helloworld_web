package main

import (
  "net/http"
  "log"
  "github.com/gorilla/mux"
  "helloworld_web/app/db"
  "helloworld_web/app/route"
  "helloworld_web/app/handler"
)

func init(){
  db.InitDB()
}

func main() {
  r := mux.NewRouter()
  r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
  r.Handle("/", handler.HandleHome).Methods("GET")
  r.Handle("/register", handler.HandleUserNew).Methods("GET")
  r.Handle("/register", handler.HandleUserCreate).Methods("POST")
  r.Handle("/login", handler.HandleSessionNew).Methods("GET")
  r.Handle("/login", handler.HandleSessionCreate).Methods("POST")
  r.Handle("/sign-out", handler.HandleSessionDestroy).Methods("GET")
  r.Handle("/products", route.AuthMiddleware(handler.HandleProducts)).Methods("GET")
  r.Handle("/products/new", route.AuthMiddleware(handler.HandleProductNew)).Methods("GET")
  r.Handle("/products/new", route.AuthMiddleware(handler.HandleProductCreate)).Methods("POST")
  log.Fatal(http.ListenAndServe(":3000", r))
}
