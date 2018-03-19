package handler

import (
  "net/http"
  "helloworld_web/app/store"
  "helloworld_web/app/model"
  "helloworld_web/app/util"
  "helloworld_web/app/route"
)

var HandleUserNew = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  route.RenderTemplate(w, r, "users/new", nil)
})

var HandleUserCreate = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  // Process creating a user
  user, err := model.NewUser(
    r.FormValue("username"),
    r.FormValue("email"),
    r.FormValue("password"),
  )

  if(err == nil){
    existingUser, _ := store.GlobalUserStore.FindByUsernameOrEmail(user.Username, user.Email)
    if existingUser.ID != "" {
      err = util.ErrUsernameOrEmailExists
    }
  }

  if err != nil {
    if util.IsValidationError(err) {
      route.RenderTemplate(w, r, "users/new", map[string] interface{}{
        "Error": err.Error(),
        "User": user,
      })
      return
    }
  }

  err = store.GlobalUserStore.Save(&user)
  if err != nil {
    panic(err)
  }

  cookie    :=    http.Cookie{ Name: "Token", Value: route.GenerateToken(user.Username)}
  http.SetCookie(w, &cookie)
  http.Redirect(w, r, "products", http.StatusFound)
})
