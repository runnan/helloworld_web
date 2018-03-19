package handler

import (
  "net/http"
  "strconv"
  "helloworld_web/app/store"
  "helloworld_web/app/model"
  "helloworld_web/app/route"
)

var HandleProducts = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  products, err := store.GlobalProductStore.FindAll(0)
  if err != nil {
    panic(err)
  }
  route.RenderTemplate(w, r, "products/index", map[string]interface{}{
    "Products": products,
  })
})

var HandleProductNew = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  route.RenderTemplate(w, r, "products/new", nil)
})

var HandleProductCreate = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  product := model.NewProduct()
  product.Description = r.FormValue("description")
  product.Name = r.FormValue("name")
  f, _ := strconv.ParseFloat(r.FormValue("price"), 64)
  product.Price = f

  store.GlobalProductStore.Save(product)
  http.Redirect(w, r, "/products", http.StatusFound)
})
