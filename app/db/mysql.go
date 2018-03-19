package db

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "helloworld_web/app/util"
)

var GlobalMySQLDB *sql.DB

func InitDB() {
  db, err := NewMySQLDB(util.GetConfigurationValue("mydb"))
  if err != nil {
    panic(err)
  }
  GlobalMySQLDB = db
}

func NewMySQLDB(dsn string) (*sql.DB, error) {
  db, err := sql.Open("mysql", dsn+"?parseTime=true")
  if err != nil {
    return nil, err
  }
  return db, db.Ping()
}
