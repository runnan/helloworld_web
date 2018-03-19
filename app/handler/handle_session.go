package handler

import (
  "net/http"
  "helloworld_web/app/store"
  "helloworld_web/app/util"
  "helloworld_web/app/route"
)

var HandleSessionDestroy = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  cookie    :=    http.Cookie{ Name: "Token", MaxAge: -1 }
  http.SetCookie(w, &cookie)
  http.Redirect(w, r, "/login", http.StatusFound)
})

var HandleSessionNew = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  next := r.URL.Query().Get("next")
  route.RenderTemplate(w, r, "sessions/new", map[string]interface{}{
    "Next": next,
  })
})

var HandleSessionCreate = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  username := r.FormValue("username")
  password := r.FormValue("password")
  next := r.FormValue("next")
  user, err := store.GlobalUserStore.FindUser(username, password)
  if err != nil {
    if util.IsValidationError(err) {
      route.RenderTemplate(w, r, "sessions/new", map[string]interface{}{
        "Error": err,
        "User": user,
        "Next": next,
      })
      return
    }
    panic(err)
  }
  cookie    :=    http.Cookie{ Name: "Token", Value: route.GenerateToken(username)}
  http.SetCookie(w, &cookie)
  http.Redirect(w, r, "products", http.StatusFound)
})
