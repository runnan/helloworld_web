package model
import(
  "github.com/leekchan/accounting"
  "helloworld_web/app/util"
)

type Product struct{
  ID string
  Name string
  Price float64
  Description string
}

func NewProduct() *Product{
  return &Product{
    ID: util.GenerateID("P", 10),
  }
}

func (product *Product) GetPrice() string {
  ac := accounting.Accounting{Symbol: "VND "}
  return ac.FormatMoney(product.Price)
}
