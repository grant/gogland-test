package main

import (
  "fmt"
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/", hello)
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }
  fmt.Println("listening", port)
  err := http.ListenAndServe(":5000", nil)
  if err != nil {
    panic(err)
  }
}

func hello(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(res, "hello, world!")
}