package main


import (
  "demo-tmdb-proxy-api/handlers"
  "net/http"
  "log"
  "os"
)

func main() {
  http.HandleFunc("/", handlers.ProxyHandler)

  log.Println("Proxy server running")

  port := os.Getenv("PORT")
  if port == "" {
    port = "3080"
  }

  err := http.ListenAndServe(":" + port, nil)
  if err != nil {
  	log.Fatal(err)
  }
}
