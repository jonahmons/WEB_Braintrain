package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("html"))
  http.Handle("/html/", http.StripPrefix("/html/", fs))

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}
