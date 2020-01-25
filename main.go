package main

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

const DRIVER = "mysql"
const DSN = "root:password@tcp(mysql:3306)/matsun_dev"


func main() {
  db, err := sql.Open(DRIVER, DSN)
  if err != nil {
      fmt.Println("Openエラー")
  } else {
    fmt.Println("Open OK!")
  }

  err = db.Ping()
  if err != nil {
    fmt.Println("接続失敗")
  } else {
    fmt.Println("接続OK")
  }

  db.Close()


}


