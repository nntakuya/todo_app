package main

import (
  "github.com/nntakuya/todo_app/db"
  "github.com/nntakuya/todo_app/server"
)

func main() {
  db.Init()

  server.Init()
}


