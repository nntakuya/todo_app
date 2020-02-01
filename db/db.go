package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/nntakuya/todo_app/entity"
)

var (
  db *gorm.DB
  err error
)

// Init is initialize db from mainfunction
func Init() {
  db, err = gorm.Open("mysql", "root:password@tcp(docker.for.mac.localhost:3306)/todo_dev")

  if err != nil {
    panic(err)
  }

  autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
  return db
}


// Close is closing db
func Close() {
  if err := db.Close(); err != nil {
    panic(err)
  }
}

func autoMigration() {
  db.AutoMigrate(&entity.User{})
}

