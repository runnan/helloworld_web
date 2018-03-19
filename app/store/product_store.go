package store
import (
  // "database/sql"
  "helloworld_web/app/model"
  "helloworld_web/app/db"
)

var GlobalProductStore ProductStore
const pageSize = 25

type ProductStore struct{

}

func (store *ProductStore) Save(product *model.Product) error {
  _,
  err := db.GlobalMySQLDB.Exec(`REPLACE INTO products (id, name, description, price) VALUES (?, ?, ?, ?)`,
  product.ID,
  product.Name,
  product.Description,
  product.Price,
)
return err
}

func (store *ProductStore) FindAll(offset int) ([]model.Product, error) {
  rows, err := db.GlobalMySQLDB.Query(
    `SELECT id, name, description, price
    FROM products
    LIMIT ?
    OFFSET ?`,
    pageSize,
    offset,
  )
  if err != nil {
    return nil, err
  }
  products := []model.Product{}
  for rows.Next() {
    product := model.Product{}
    err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price,)
    if err != nil {
      return nil, err
    }
    products = append(products, product)
  }
  return products, nil
}
