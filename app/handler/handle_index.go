package handler

import (
  "net/http"
)

var HandleHome = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "/products", http.StatusFound)
})
